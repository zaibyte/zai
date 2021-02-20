package client

import (
	"io"
	"time"
)

// Client is Zai client.
// It shouldn't be used after calling Close().
type Client interface {
	// PutObj puts object to the box which this Zai belongs to.
	// Returns oid & any error.
	PutObj(objData io.Reader, size int64, timeout time.Duration) (oid uint64, err error)
	// GetObj gets object by oid.
	// start & end is the range bytes of object: [start, end)
	// Both of start & end will be aligned to uid.GrainSize.
	// The min range is 1 uid.GrainSize, when both of start & end are 0, means getting
	// the whole object.
	// If end is beyond the object length, it will get all left bytes.
	GetObj(oid uint64, objData io.Writer, start, end int64, timeout time.Duration) error
	// DeleteObj deletes object by oid.
	DeleteObj(oid uint64, timeout time.Duration) error
	// UpdateObj updates oid's data [start, end) with newData.
	// Returns newOid, after updating.
	UpdateObj(oid uint64, start, end int64, newData io.Reader) (newOid uint64, err error)
	// Close closes the client.
	Close()
}
