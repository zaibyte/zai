package client

import (
	"time"
)

const (
	// These timeout is per RPC request timeout.
	// If there is retry, the final cost maybe higher than it (if need retry)
	//
	// 3 seconds is much more than enough for Client request,
	// because the max transport size is 4MB.
	DefaultPutTimeout = 3 * time.Second
	DefaultGetTimeout = 3 * time.Second
	DefaultDeleteTimeout = 3 * time.Second
	DefaultDeleteBatchTimeout = 3 * time.Second
)

// Client is Zai client.
// It shouldn't be used after calling Close().
type Client interface {
	// PutObj puts object to the box which this Zai belongs to.
	// Returns non-zero oid if err is nil.
	// PutObj will read the whole objData and return its length.
	//
	// PutObj will try to put object to a ground until succeed or reach the max retires.
	PutObj(objData []byte, timeout time.Duration) (oid uint64, size int64, err error)
	// GetObj gets object by oid.
	// range bytes of object: [offset, offset+n]
	// If the offset < 0, start with 0.
	// If offset+n > object_size, it will read until meeting the end of the object.
	// n is the length you want to get.
	// Return how many bytes written into buf.
	// GetObj will try to connect all replicas until succeed or reach the max retires.
	// It will return orpc.ErrReplicasCollapse if all replicas return disk/extent_broken/ghost/clone error.
	GetObj(oid uint64, buf []byte, offset, n int64, isClone bool, timeout time.Duration) (size int64, err error)
	// DeleteObj deletes object by oid.
	DeleteObj(oid uint64, timeout time.Duration) error
	// DeleteBatch deletes multi objects in a single RPC call.
	DeleteBatch(oids []uint64, timeout time.Duration) error

	// Start starts Client.
	Start() error
	// Stop stops Client, release resource.
	Stop()
	// Close closes Client with an error which will be passed to the pending requests.
	Close(err error)
}
