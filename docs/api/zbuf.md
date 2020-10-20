ZBuf API
=====

ZBuf is Zai's replicas storage server. The core features' protocol is Zai's rpc(based on TCP).

## Core features

	PutObj(reqid uint64, oid [16]byte, objData xbytes.Buffer) error
    GetObj(reqid uint64, oid [16]byte) (objData xbytes.Buffer, err error)
    DeleteObj(reqid uint64, oid [16]byte) error

