Design Logic
===

*"There is no secret in source code."*

Buf if we don't have a clear logic, we'll get lost in details.

Here I will explain the internal logic of Zai.

## Question and Target

Zai is designed as Tesamc Machine Learning (mainly) storage system. The total capacity
is not large, but the I/O through & parallel pressure is high in present. We need a low-cost
storage system to provides enough ability of I/O with few nodes and disks:

1. High space utilization

2. Big data blocks parallel I/O supports

These are the core abilities which Zai must have.

We need this too:

> Simple Maintenance

The cost of scale should be low. The codes of project should be easy to modify/upgrade,
because we often need to customized storage system for satisfying the real needs.

## Design Process

Let's begin to make decisions from top to low level.

### Multi-Zone/Region Supports

Zai doesn't implement the geographical disaster recovery directly. It could be built on the top
of Zai, and it's up to users.

### High Availability

Not single point in Zai.

Entrance layer is stateless. Manager Services are built on consensus algorithm.

### High Durability

Zai uses copies and Erasure Codes to provide durability. There is also checksum in storage media, and E2E checking is enabled by default.

### High Performance

Performance is linear with nodes number. In single node, performance is mainly depends on hardware. The most important thing Zai should do is implementing a light I/O stack.

The next important thing is balance, avoiding too many hot spots scheduled in the same nodes/disk.

### Large Cluster

Cluster in Zai is made of multi boxes. Box is a kind of sub-cluster. Each box is individual in logic.

Sub-cluster makes things clear and easy to maintain.

### Monitor

Zai provides status API, and collects Prometheus metrics.

Status API is using for showing basic information.

Metrics is using for monitor cluster detailed status and performance information.

### File System Protocol Supports

File system could be built on the top of Zai.

There are companies are working on it. (e.g. [JuiceFS](https://juicefs.com))

### Indexes

There are multi indexes levels in Zai for reducing index space overhead.

First level is logical position in oid, second level is physical offset on storage nodes, at the same time, there are records to map logical position to physical media.

More details about Index, [here](index.md)

### GC

A fast GC is implemented in Zai write buffer.

There will be not much extra space need in GC process and fast, it's important for business model which deletes files often. 

### User Defined Storage

Zai is not only a out of box storage system. It should be easily modified for professional users to fit needs better. The simple design and detailed guide are important.

#### Entrance SDK

Entrance is not just an application could be started individually, it could be embedded in users' application entrance layer too, users could access Zai directly.

#### Buffer Interface

Write buffer is a performance & cost sensitive layer. Zai's default implementations maybe not the best choices for all users, what users need to do is just satisfying buffer's interface.

#### Full Documents about Design & Development

There are full documents explain why and how, step by step to show users to build their own system.
