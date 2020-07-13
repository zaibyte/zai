Optimization
===

## I/O

**1. Set the I/O scheduler to "noop"**

Does not make re-orderings that can compete or clash with the underlying
intelligent I/O controllers. It's a good choice for both SSD and SATA
, because modern SATA also has NCQ too

**2. Set nr_requests to 1024 for HDD**

Help to sort random I/O to NCQ(saving seek time),
and avoiding stall for some call, e.g. SyncFileRange (async mode).

But there is no need to change the default value for SSD, cuz
on SSD there is no disk seek

e.g.

nr_requests = 1024

max_sectors_kb = 128KB

queue_size = 2 * 128KB * 1024 = 256MB

if nr_requests is default value (128)

queue_size = 2 * 12KB * 128 = 32MB
