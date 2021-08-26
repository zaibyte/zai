package client

import (
	"sync"
	"time"

	"g.tesamc.com/IT/keeper/client"
	"g.tesamc.com/IT/zaipkg/linkobj"
	"g.tesamc.com/IT/zaipkg/orpc"
	"g.tesamc.com/IT/zaipkg/uid"
	"g.tesamc.com/IT/zaipkg/xdigest"
	"g.tesamc.com/IT/zaipkg/xerrors"
	"g.tesamc.com/IT/zproto/pkg/keeperpb"
)

// memZai is a Zai Client built for testing purpose
// which using memory for mocking Zai Box features.
type memZai struct {
	sync.RWMutex
	// boxID & groupID here for building oid,
	// memZai is only made for ext.v1 testing, it's okay to share the same groupID.
	boxID   uint32
	groupID uint32
	OidData map[uint64][]byte
}

var _ Client = new(memZai)

func NewMemZai() *memZai {
	mz := new(memZai)
	mz.OidData = make(map[uint64][]byte)
	mz.boxID = 1
	mz.groupID = 1
	return mz
}

func (m *memZai) PutObj(objData []byte, digest uint32, timeout time.Duration) (oid uint64, err error) {
	m.Lock()
	defer m.Unlock()

	return m.put(objData, digest, uid.NormalObj), nil
}

func (m *memZai) put(data []byte, digest uint32, otype uint8) uint64 {
	oid := uid.MakeOID(m.boxID, m.groupID, uint32(len(data)/uid.GrainSize), digest, otype)

	m.OidData[oid] = data
	return oid
}

func (m *memZai) MakeLink(oids []uint64, timeout time.Duration) (oid uint64, err error) {

	if len(oids) > linkobj.MaxObjsInLink {
		return 0, xerrors.WithMessage(orpc.ErrBadRequest, "too many objects for link_obj")
	}

	data := make([]byte, linkobj.CalcLen(int64(len(oids))))

	linkobj.Make(oids, data)

	d := xdigest.Acquire()
	defer xdigest.Release(d)

	_, _ = d.Write(data)
	di := d.Sum32()

	oid = m.put(data, di, uid.LinkObj)

	// For easy implement GetObj, combine all objects as a link.
	var linkON int
	for _, id := range oids {
		linkON += int(uid.GetGrains(id)) * uid.GrainSize
	}
	linkO := make([]byte, linkON)
	start := 0
	for _, id := range oids {
		on := int(uid.GetGrains(id)) * uid.GrainSize
		copy(linkO[start:start+on], m.OidData[id])
		start += on
	}

	m.OidData[oid] = linkO

	return
}

func (m *memZai) GetObj(oid uint64, buf []byte, offset, n uint64, isClone bool, timeout time.Duration) (err error) {
	m.RLock()
	defer m.RUnlock()

	d, ok := m.OidData[oid]
	if !ok {
		return orpc.ErrNotFound
	}
	if n == 0 {
		n = uint64(len(d))
	}

	copy(buf[offset:offset+n], d)

	return nil
}

func (m *memZai) GetKeeperClient() keeperpb.KEEPERClient {
	return nil
}

func (m *memZai) GetFastKeeper() client.FClient {
	return nil
}

func (m *memZai) DeleteObj(oid uint64, timeout time.Duration) error {
	m.Lock()
	defer m.Unlock()

	delete(m.OidData, oid)
	return nil
}

func (m *memZai) Close(err error) {
	return
}
