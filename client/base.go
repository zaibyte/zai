package client

import (
	"io"
	"time"

	"g.tesamc.com/IT/keeper/client"
)

// BaseClient is the basic ObjClient implementation without cache or other advanced features.
type BaseClient struct {
}

var _baseClient ObjClient = new(BaseClient)

func NewBaseClient(kc client.FClient) *BaseClient {

}

func (b *BaseClient) UpdateObj(oid uint64, start, end int64, newData io.Reader) (newOid uint64, err error) {
	panic("implement me")
}

func (b *BaseClient) PutObj(objData []byte, timeout time.Duration) (oid uint64, size int64, err error) {
	panic("implement me")
}

func (b *BaseClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (size int64, err error) {
	panic("implement me")
}

func (b *BaseClient) DeleteObj(oid uint64, timeout time.Duration) error {
	panic("implement me")
}

func (b *BaseClient) Start() error {
	panic("implement me")
}

func (b *BaseClient) Stop() {
	panic("implement me")
}

func (b *BaseClient) Close(err error) {
	panic("implement me")
}
