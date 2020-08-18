Design Logic
===

*"There is no secrets in source code."*

Buf if we don't have a clear logic, we'll get lost in details.

Here I will explain the internal logic of Zai.

## Question and Target

LOSF(Lots of small files) is a classic question in distributed storage system.
Many people have made efforts on it, and there are many excellent projects too.

But I think the key is not just technique or engineering, 
the real problem which most people ignore is the business model.

### Classical Business Model

#### SNS

> Instagram

Number of Photos & Videos uploaded per day: [100 million+](https://www.omnicoreagency.com/instagram-statistics/)

> Tencent

Number of files uploaded per day: [several billions,and 50% files will be deleted in 90days](https://cloud.tencent.com/developer/article/1020435)

**Needs:**

It's not just a thing about speed.

Data could get cold soon in SNS business model, and the the amount of small files is crazy. So we need to get balance of performance and cost.

What's more, users Do delete a lot (and SNS companies wouldn't like to keep all data forever either, they will delete media files in chatting records automatically if users don't do the saving operation). But in most implementations, we often combine small files as a big one for reducing
the inode expense, but it cause Garbage Collection problem. The traditional way of GC is too slow to satisfy their needs.

#### Video/Voice surveillance

HLS is a common tech in video surveillance, and millions clients run 7x24hours, so there will be lots of small files.

And usually we won't keep surveillance data forever, after 7-45days we will delete it. It needs a 
fast GC too.

#### Machine Learning

Actually, ML is some kind of offline business. What they want is good performance when
they need the data, and low cost when they don't need it.

The gap between computer node and storage node is high, we usually use multi-levels cache in 
storage system which supports machine learning jobs.

#### CDN

More and more data will be put in CDN for accelerating network performance for avoiding backend servers overload. But in present, most of CDNs are just using rough file systems to provide cache
services. What they need is a storage system to support various request, high IPOS and low latency.
What's more, they need this storage system could have a efficient GC. 

### They Do Want These

> Data Tiers

Data could be transferred to cheaper media automatically when they getting cold.

As previously mentioned, we always append small files to a big file in LOSF solution.
So it becomes harder to only transfer the cold one, because it will leave lots of "holes"
in the big file. Some public cloud companies regard the big file as the smallest unit
in the transfer process, so parts of hot files will be regraded as "cold" too or too many cold
files will be regarded as "hot".

> GC in short time

It's better to save disk I/O in GC process. At the same time, it should be fast,
because there may be too many delete operations. And data tires need high perf GC too.

> Simple Maintenance

The cost of scale should be low. The codes of project should be easy to modify/upgrade,
because they often need to customized their own storage system for satisfying the real needs.

Center-less distributed storage system is cool, but hard to scale, the rebalance is a nightmare. 
And they often regard user-defined metadata as a special object, and even tens of million
objects in a same bucket may cause break down. Or they just put these metadata in a non-distributed
database(sure, it will have some methods to improve availability, but not enough), weak and dangerous.

And some systems are too ambitious, S3 API, file API, block API... Yes, it's powerful and useful.
But it's too much for a company which just want to put/get billion objects. How can they build
their own system? What about the cost of hiring a team to maintain the system? 

## Design Process

Let's begin to make decisions from top to low level.

### Multi-Zone/Region Supports

Zai doesn't implement the geographical disaster recovery directly. It could be built on the top
of Zai, and it's up to users.

Most of data will be transferred to Erasure Codes storage system, the cost isn't that high. 
And it's fair enough that pay more if you want extra features.

### High Availability

Not single point in Zai.

Entrance layer is stateless. Manager Services are built on consensus algorithm.
Data has three copies/split into several pieces (Erasure Codes) and any one of them could answer read request.

### High Durability

Zai uses copies and Erasure Codes to provide durability. There is also checksum in storage media, and E2E checking is enabled by default.

More discussion about durability is in [durability](../conversation-street/durability.md).

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

### S3 API Supports

Zai doesn't support S3 naturally. It could be built on the top of Zai.

### File System Protocol Supports

File system could be built on the top of Zai.

There are companies are working on it. (e.g. [JuiceFS](https://juicefs.com))

### Block Device Supports

Block Device will not be provided in future. 

Keep I/O stack clean.

### Stream Storage Supports

Stream Storage could be deployed on the top of Zai.

Zai provides the ability of accessing historical data.

### Various Metadata Supports

Zai doesn't support various metadata. It could be built on the top of Zai.

The main reason is that users may have their own metadata service already for supporting their
business models. What they will do is only add a field to store file's [oid](basic_concepts.md) if we want to use Zai. 

### Data Tiers

Zai assumes 90% data is warm/cold should be putted into cheaper devices. In most cases, 
data getting cold with time pass away.

It'll be great if Zai cares about more details of data's "temperature", but it's too complex and
not easy to implement. And we could have extra cache layer(optional) if the scheduler strategy is not
enough.

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
