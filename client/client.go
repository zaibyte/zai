package client

import (
	"io"
	"time"
)

// Client is Zai client.
// It shouldn't be used after calling Close().
type Client interface {
	// PutObj puts object to the box which this Zai belongs to.
	// Returns non-zero oid if err is nil.
	// PutObj will read objData until io.EOF, and return the read.
	PutObj(objData io.Reader, timeout time.Duration) (oid uint64, read int64, err error)
	// GetObj gets object by oid.
	// range bytes of object: [offset, offset+n]
	// If the offset < 0, start with 0.
	// If offset+n > object_size, it will read until meeting the end of the object.
	// Return how many bytes written into objData.
	// GetObj will try to connect all replicas untill succeed.
	// It will return ErrReplicasCollapse if all replicas return disk/extent_broken/ghost/clone error.
	GetObj(oid uint64, objData io.Writer, offset, n int64, isClone bool, timeout time.Duration) (written int64, err error)
	// DeleteObj deletes object by oid.
	DeleteObj(oid uint64, timeout time.Duration) error
	// UpdateObj updates oid's data [offset, offset+newData_len] with newData.
	// Returns newOid and newData_len, after updating.
	UpdateObj(oid uint64, offset, newData io.Reader) (newOid uint64, read int64, err error)
	// Close closes the client.
	Close()
}
