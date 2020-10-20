ZMap API
=====

ZMap will be embedded in Keeper.

## Core features

	// Put puts key:oid pair to ZMap.
	Put(oid, key string) error
	
	// Get gets key's oid if has.
	Get(key string) (oid string, err error)
			
	// Delete deletes key:oid pair.
	Delete(key string) error

