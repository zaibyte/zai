Consistency
===

## Objects Operation

There are only three main operations with objects you can take in Zai:

1. Put
2. Get
3. Delete

So consistency issues are bit easier in Zai. But still could cause a disaster if don't take care.

### Put

There is only one choice of quorum in Zai for put:

Only if all replicas success, return OK. 

### Get

Any extent could response get request.

### Delete

Delete operation is async. Return success immediately.

Unlike ``Get`` , delete can't choice another group when meet failure, and delete won't cause
collapse, so for improving availability using async way to process delete is a not bad choice.

## Extent Operation

The major issue is repairing extent.

After sealing, chose the shortest extent as the source.

### Dirty Data Issue

Dirty Data may happen when delete operations are not consistent within replicas,
choice the shortest extent when repairing needed may help but still can't fix the issue.

It can't avoid in present mode. It's rare, but it Do happen.

## Cluster Manager

Keeper(managing component) embed raft consensus algorithm for providing strong
consistency data service.



