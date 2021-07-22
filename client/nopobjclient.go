package client

import (
	"time"

	kc "g.tesamc.com/IT/keeper/client"
	"g.tesamc.com/IT/zproto/pkg/keeperpb"
)

// NopObjClient implements ObjClient but does nothing for testing only.
type NopObjClient struct {
}

var _nopObjClient ObjClient = NopObjClient{}

func (c NopObjClient) PutObj(objData []byte, digest uint32, timeout time.Duration) (oid uint64, err error) {

	return 0, nil
}

func (c NopObjClient) MakeLink(oids []uint64, timeout time.Duration) (oid uint64, err error) {
	return 0, nil
}

func (c NopObjClient) GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (err error) {
	return nil
}

func (c NopObjClient) DeleteObj(oid uint64, timeout time.Duration) error {
	return nil
}

func (c NopObjClient) GetKeeperClient() keeperpb.KEEPERClient {
	return nil
}

func (c NopObjClient) GetFastKeeper() kc.FClient {
	return nil
}

func (c NopObjClient) Close(err error) {
	return
}
