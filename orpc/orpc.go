package orpc

import (
	"time"

	"g.tesamc.com/IT/zaipkg/orpc"
)

// OClient provides orpc.Client pool.
type OClient struct {
}

func New() orpc.Client {

}

func (c *OClient) Start() error {
	panic("implement me")
}

func (c *OClient) Stop() {
	panic("implement me")
}

func (c *OClient) Close(err error) {
	panic("implement me")
}

func (c *OClient) PutObj(reqid uint64, oid uint64, extID uint32, objData []byte, timeout time.Duration) error {
	panic("implement me")
}

func (c *OClient) GetObj(reqid uint64, oid uint64, extID uint32, objData []byte, isClone bool, timeout time.Duration) error {
	panic("implement me")
}

func (c *OClient) DeleteObj(reqid uint64, oid uint64, extID uint32, timeout time.Duration) error {
	panic("implement me")
}

func (c *OClient) DeleteBatch(reqid uint64, oids []uint64, extID uint32, timeout time.Duration) error {
	panic("implement me")
}
