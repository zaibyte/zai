package client

import (
	"context"
	"time"

	"g.tesamc.com/IT/zaipkg/xdigest"

	kc "g.tesamc.com/IT/keeper/client"
)

// ProClient is professional ObjClient which implements ObjClient interface with cache & advanced features.
type ProClient struct {
	cfg *ProClientConfig
	kfc kc.FClient
}

var _proClient ObjClient = new(ProClient)

// ProClientConfig is the configs of ProClient.
type ProClientConfig struct {
	KeeperConfig kc.FCConfig `toml:"keeper_config"`
}

func NewProClient(cfg *ProClientConfig) (*ProClient, error) {

	kfc, err := kc.NewFC(cfg.KeeperConfig)
	if err != nil {
		return nil, err
	}

	return &ProClient{
		cfg: cfg,
		kfc: kfc,
	}, nil
}

func (p *ProClient) PutObj(objData []byte, digest uint32, timeout time.Duration) (oid uint64, size int64, err error) {

	if digest == 0 {
		xdigest.Sum32(objData)
	}

	for i := 0; i < 3; i++ {

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		p.kfc.GetWritableGroup(ctx, len(objData))
	}
}

func (p *ProClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (size int64, err error) {
	panic("implement me")
}

func (p *ProClient) DeleteObj(oid uint64, timeout time.Duration) error {
	panic("implement me")
}

func (p *ProClient) Start() error {
	panic("implement me")
}

func (p *ProClient) Stop() {
	panic("implement me")
}

func (p *ProClient) Close(err error) {
	panic("implement me")
}
