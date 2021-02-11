package client

import (
	"io"
	"time"
)

// BaseClient is the basic Client implementation without cache or other advanced features.
type BaseClient struct {

}

func (b *BaseClient) UpdateObj(oid uint64, start, end int64, newData io.Reader) (newOid uint64, err error) {
	panic("implement me")
}

func (b *BaseClient) PutObj(objData io.Reader, size int64, timeout time.Duration) (oid uint64, err error) {

	panic("implement me")
}

func (b *BaseClient) GetObj(oid uint64, objData io.Writer, start, end int64, timeout time.Duration) error {
	panic("implement me")
}

func (b *BaseClient) DeleteObj(oid uint64, timeout time.Duration) error {
	panic("implement me")
}

func (b *BaseClient) Close() {
	panic("implement me")
}

