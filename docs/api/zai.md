Zai API
=====

Zai will be embedded in other applications (SDK), and has an HTTP server which provides control entrance.

## Core features

	// PutObj puts object to the ZBuf node which Objecter connected.
	PutObj(reqid uint64, oid string, objData []byte, timeout time.Duration) error
	
	// GetObj gets object from the ZBuf node which Objecter connected.
	GetObj(reqid uint64, oid string, timeout time.Duration) (obj io.ReadCloser, err error)
		
	// DeleteObj deletes object in the ZBuf node which Objecter connected.
	DeleteObj(reqid uint64, oid string, timeout time.Duration) error

