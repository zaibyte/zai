package client

import (
	"time"

	kc "g.tesamc.com/IT/keeper/client"
	"g.tesamc.com/IT/zproto/pkg/keeperpb"
)

// These timeout is per RPC request timeout.
// If there is retry, the final cost maybe higher than it (if need retry)
//
// 3 seconds is much more than enough for ObjClient request,
// because the max transport size is 4MB.
const (
	DefaultMaxPutRetries      = 3 // There is no need to retry many times.
	DefaultMaxGetRetries      = 3
	DefaultPutTimeout         = 3 * time.Second
	DefaultGetTimeout         = 3 * time.Second
	DefaultDeleteTimeout      = 3 * time.Second
	DefaultDeleteBatchTimeout = 3 * time.Second
)

// Client is Zai client.
// It shouldn't be used after calling Close().
type Client interface {
	ObjClient

	// Close closes ObjClient with an error which will be passed to the pending requests.
	Close(err error)
}

// ObjClient is Zai client which provides uid.NormalObj object(in byte slice) accessing.
// ObjClient could use xbytes.GetBytes(n int) as buffer for saving cost of memory allocation & GC.
type ObjClient interface {
	// PutObj puts object to the box which this Zai belongs to (relies on which keeper it's using).
	// Returns non-zero oid if err is nil.
	//
	// objData's size must already be aligned to uid.GrainSize.
	// PutObj will read the whole objData.
	//
	// PutObj will try to put object to a group until succeed or reach the max retires.
	//
	// digest is the digest of objData.
	PutObj(objData []byte, digest uint32, timeout time.Duration) (oid uint64, err error)
	// GetObj gets object by oid.
	// range bytes of object: [offset, offset+n)
	// n is the length you want to get.
	// GetObj will try to connect all replicas until succeed.
	//
	// Warn:
	// We assume len(buf) = n,
	// offset & n are legal. Both of offset & n are in object size.
	// If n == oid.grains * grain_size, ignore offset.
	GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (err error)
	// DeleteObj deletes object by oid.
	DeleteObj(oid uint64, timeout time.Duration) error

	// GetKeeperClient gets keeperpb.KEEPERClient for invokers.
	// There must be keeper client inside ObjClient.
	GetKeeperClient() keeperpb.KEEPERClient

	// GetFastKeeper returns inside kc.FClient for invokers.
	GetFastKeeper() kc.FClient

	Close(err error)
}
