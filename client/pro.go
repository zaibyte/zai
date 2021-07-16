package client

import (
	"time"
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
	return p.base.PutObj(objData, digest, timeout)
}

func (p *ProClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (err error) {
	panic("implement me")
}

func (p *ProClient) DeleteObj(oid uint64, timeout time.Duration) error {
	return p.base.DeleteObj(oid, timeout)
}

func (p *ProClient) Close(err error) {
	p.base.Close(err)
}
