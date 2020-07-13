Basic Concepts
==

### oid

The unique ID of an object.       

>- Type: string

### extent

The collection of objects on local storage node's filesystem.   
Every extent belong to a group.

>- Type: Continuous space on disk

#### extentID

extentID is the extent's name.

>- extentID Type: string, `<group_id>-<seq>`, e.g. `"12345-001"`
>- seq Type: uint8

### group
The unit of fault-tolerance of extents, it could be 3 replicas of one extent or an Erasure Code Stripe of several extents.

Any extent in the same group can be recovered by others.

#### groupID

And each group has an unique ID, groupID.

>- Type: uint32
   
### diskID

Every disk in Zai must have an unique diskID.

>- Type: UUID

### node

Machine which being deployed Zai services.

#### nodeID

Every ZBuf/ZCold component node must have an unique ID. It's node's MAC address in hex by default.

>- Type: string

### box 

A logical independent Zai cluster, every box has a boxID.

#### boxID

>- Type: uint32 (`[1,2^24)`)
   
### namespace

Multi boxes could be regard as a namespace if their boxIDs are unique.

Because every oid has a boxID in it, so it's easy to write a router to indicate the data flows.
   
