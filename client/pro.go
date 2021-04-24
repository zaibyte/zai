package client

import (
	"io"
	"time"
)

// ProClient is professional Client which implements Client interface with cache & advanced features.
// It's the Client in production env.
type ProClient struct {
	cfg *ProClientConfig `toml:"cfg"`
}

// ProClientConfig is the configs of ProClient.
type ProClientConfig struct {
	BoxConfigs []BoxConfig `toml:"box_configs"`
}

// BoxConfig is the configs of ProClient for each box.
type BoxConfig struct {
	BoxID       uint32   `toml:"box_id"`
	KeeperAddrs []string `toml:"keeper_addrs"`
}


func (p *ProClient) PutObj(objData io.Reader, timeout time.Duration) (oid uint64, read int64, err error) {
	panic("implement me")
}

func (p *ProClient) GetObj(oid uint64, objData io.Writer, offset, n int64, isClone bool, timeout time.Duration) (written int64, err error) {
	panic("implement me")
}

func (p *ProClient) DeleteObj(oid uint64, timeout time.Duration) error {
	panic("implement me")
}

func (p *ProClient) UpdateObj(oid uint64, offset, newData io.Reader) (newOid uint64, read int64, err error) {
	panic("implement me")
}

func (p *ProClient) Close() {
	panic("implement me")
}

