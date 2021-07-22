package client

import (
	"time"

	"g.tesamc.com/IT/zaipkg/linkobj"
	"g.tesamc.com/IT/zaipkg/orpc"
	"g.tesamc.com/IT/zaipkg/xerrors"

	"g.tesamc.com/IT/zaipkg/uid"
	"g.tesamc.com/IT/zaipkg/xbytes"
	"g.tesamc.com/IT/zaipkg/xdigest"

	kc "g.tesamc.com/IT/keeper/client"

	"g.tesamc.com/IT/zproto/pkg/keeperpb"
)

// ProClient is professional ObjClient which implements ObjClient interface with cache & advanced features.
type ProClient struct {
	base *BaseClient
}

var _proClient Client = new(ProClient)

func NewProClient(cfg *Config) (*ProClient, error) {

	b, err := NewBaseClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ProClient{
		base: b,
	}, nil
}

func (p *ProClient) PutObj(objData []byte, digest uint32, timeout time.Duration) (oid uint64, err error) {
	return p.base.PutObj(objData, digest, uid.NormalObj, timeout)
}

func (p *ProClient) MakeLink(oids []uint64, timeout time.Duration) (oid uint64, err error) {

	if len(oids) > linkobj.MaxObjsInLink {
		return 0, xerrors.WithMessage(orpc.ErrBadRequest, "too many objects for link_obj")
	}

	data := xbytes.GetBytes(int(linkobj.CalcLen(int64(len(oids)))))
	defer xbytes.PutBytes(data)

	linkobj.Make(oids, data)

	d := xdigest.Acquire()
	defer xdigest.Release(d)

	_, _ = d.Write(data)
	di := d.Sum32()

	oid, err = p.base.PutObj(data, di, uid.LinkObj, timeout)
	if err != nil {
		return 0, err
	}
	return
}

func (p *ProClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (err error) {

	otype := uid.GetOType(oid)
	if otype == uid.NormalObj {
		return p.base.GetObj(oid, buf, offset, n, isClone, timeout)
	}

	// Link Object.

}

func (p *ProClient) DeleteObj(oid uint64, timeout time.Duration) error {
	return p.base.DeleteObj(oid, timeout)
}

func (p *ProClient) GetKeeperClient() keeperpb.KEEPERClient {
	return p.base.kfc.GetClient()
}

func (p *ProClient) GetFastKeeper() kc.FClient {
	return p.base.kfc
}

func (p *ProClient) Close(err error) {
	p.base.Close(err)
}
