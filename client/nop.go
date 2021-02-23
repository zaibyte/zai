package client

import (
	"io"
	"time"
)

// NopClient implements Client interface but do nothing.
// Only for testing.
type NopClient struct {

}

func (c *NopClient) PutObj(objData io.Reader, timeout time.Duration) (oid uint64, read int64, err error) {
	return 0, 0, err
}

func (c *NopClient) GetObj(oid uint64, objData io.Writer, offset, n int64, isClone bool, timeout time.Duration) (written int64, err error) {
	return 0, err
}

func (c *NopClient) DeleteObj(oid uint64, timeout time.Duration) error {
	return nil
}

func (c *NopClient) UpdateObj(oid uint64, offset, newData io.Reader) (newOid uint64, read int64, err error) {
	return 0, 0, err
}

func (c *NopClient) Close() {
	return
}

