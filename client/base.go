package client

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"g.tesamc.com/IT/zproto/pkg/keeperpb"

	kc "g.tesamc.com/IT/keeper/client"
	"g.tesamc.com/IT/zaipkg/orpc"
	"g.tesamc.com/IT/zaipkg/orpc/otcp"
	"g.tesamc.com/IT/zaipkg/uid"
	"g.tesamc.com/IT/zaipkg/xerrors"
	"g.tesamc.com/IT/zaipkg/xlog"
	"g.tesamc.com/IT/zaipkg/xmath/xrand"
	"g.tesamc.com/IT/zproto/pkg/metapb"
)

// BaseClient is the basic client implementation for supporting Clients with cache or other advanced features.
type BaseClient struct {
	cfg      *Config
	kfc      kc.FClient
	oclients sync.Map // zBuf_addr : orpc.Client
}

func NewBaseClient(cfg *Config) (*BaseClient, error) {
	kfc, err := kc.NewFC(cfg.KeeperConfig)
	if err != nil {
		return nil, err
	}

	return &BaseClient{
		cfg: cfg,
		kfc: kfc,
	}, nil
}

func (b *BaseClient) PutObj(objData []byte, digest uint32, otype uint8, timeout time.Duration) (oid uint64, err error) {

	if timeout == 0 {
		timeout = DefaultPutTimeout
	}

	boxID := uint32(b.cfg.KeeperConfig.BoxID)
	grains := uint32(len(objData) / uid.GrainSize)
	replicas := b.kfc.GetReplicas()
	reqid := uid.MakeReqID()

	// Every time using a new exclusive, rely on background updating. It's a simple mechanism.
	// TODO may maintain a global unhealthy container in client.
	exclusive := make([]uint32, 0, 3)
	for i := 0; i < DefaultMaxPutRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		g, err2 := b.kfc.GetWritableGroup(ctx, len(objData), exclusive)
		if err2 != nil {
			cancel()
			err = err2
			continue
		}
		cancel()

		oid0 := uid.MakeOID(boxID, g.Id, grains, digest, uid.NormalObj)
		wg := new(sync.WaitGroup)
		wg.Add(replicas)

		var ok int64
		var rwDoing int

		for _, ext := range g.Exts {
			if ext.State == metapb.ExtentState_Extent_ReadWrite && rwDoing < replicas {

				go func(extid uint32) {
					defer wg.Done()

					zBuf, err2 := b.kfc.GetZBuf(context.Background(), ext.InstanceId)
					if err2 != nil {
						xlog.WarnID(reqid, err2.Error())
						return
					}

					if zBuf.State != metapb.ZBufState_ZBuf_Up {
						xlog.WarnID(reqid, "zBuf not up")
						return
					}

					// TODO in present, we only using the first address.
					c, err2 := b.getOClient(zBuf.Address[0])
					if err2 != nil {
						xlog.ErrorID(reqid, err2.Error())
						return
					}

					err2 = c.PutObj(reqid, oid0, extid, objData, timeout)
					if err2 != nil {
						xlog.ErrorID(reqid, err2.Error())
						return
					}
					atomic.AddInt64(&ok, 1)
				}(ext.Id)

				rwDoing++
			}
		}

		wg.Wait()
		if atomic.LoadInt64(&ok) < int64(replicas) {
			err = xerrors.WithMessage(orpc.ErrInternalServer, fmt.Sprintf("failed to put obj: %d to group: %d",
				oid0, g.Id))
			xlog.ErrorID(reqid, err.Error())
			exclusive = append(exclusive, g.Id)
			continue
		} else {
			err = nil
			oid = oid0
			break
		}
	}

	return
}

// GetObj gets normal object by its offset & n.
func (b *BaseClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (err error) {

	if timeout == 0 {
		timeout = DefaultGetTimeout
	}

	reqid := uid.MakeReqID()

	gid := uid.GetGroupIDFromOID(oid)
	g, err := b.kfc.GetGroup(context.Background(), gid)
	if err != nil {
		xlog.ErrorID(reqid, err.Error())
		return err
	}

	if g.State == metapb.GroupState_Group_Collapse {
		err = xerrors.WithMessage(orpc.ErrReplicasCollapse, fmt.Sprintf("failed to get obj, oid: %d", oid))
		xlog.ErrorID(reqid, err.Error())
		return err
	}

	candidates := make([]uint32, 0, len(g.Exts))
	for extID, ext := range g.Exts {
		if ext.State != metapb.ExtentState_Extent_Clone && ext.State != metapb.ExtentState_Extent_Broken {
			candidates = append(candidates, extID)
		}
	}

	if uint32(n) == uid.GetGrains(oid)*uid.GrainSize {
		offset = 0
	}

	for i := 0; i < DefaultMaxGetRetries; i++ {
		j := xrand.Uint32n(uint32(len(candidates)))
		extID := candidates[j]
		ext := g.Exts[extID]

		zBuf, err2 := b.kfc.GetZBuf(context.Background(), ext.InstanceId)
		if err2 != nil {
			xlog.WarnID(reqid, err2.Error())
			err = err2
			continue
		}

		if zBuf.State == metapb.ZBufState_ZBuf_Tombstone || zBuf.State == metapb.ZBufState_ZBuf_Disconnected {
			xlog.WarnID(reqid, "zBuf not ready for get")
			err = orpc.ErrInstanceDown
			continue
		}

		// TODO in present, we only using the first address.
		c, err2 := b.getOClient(zBuf.Address[0])
		if err2 != nil {
			xlog.WarnID(reqid, err2.Error())
			err = err2
			continue
		}

		err = c.GetObj(reqid, oid, extID, buf, offset, n, isClone, timeout)
		if err != nil {
			xlog.WarnID(reqid, err.Error())
			continue
		} else {
			return nil
		}
	}

	return
}

func (b *BaseClient) DeleteObj(oid uint64, _timeout time.Duration) error {
	return b.kfc.DeleteObj(oid)
}

func (b *BaseClient) GetKeeperClient() keeperpb.KEEPERClient {
	return b.GetKeeperClient()
}

func (b *BaseClient) GetFastKeeper() kc.FClient {
	return b.kfc
}

func (b *BaseClient) Close(err error) {

	b.kfc.Close()

	b.oclients.Range(func(key, value interface{}) bool {
		oc := value.(orpc.Client)
		oc.Close(err)
		return true
	})
}

func (b *BaseClient) getOClient(addr string) (orpc.Client, error) {

	v, ok := b.oclients.Load(addr)
	if !ok {
		c := otcp.NewClient(addr)
		c.Conns = b.cfg.ZBufConns
		err := c.Start()
		if err != nil {
			return nil, err
		}
		b.oclients.Store(addr, c)
		return c, nil
	}
	return v.(orpc.Client), nil
}
