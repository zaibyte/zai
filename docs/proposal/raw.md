RAW Proposals
===

## IO scheduler

Migrate/Repair/Get/Put/Delete have different classes.

The major problem is how many goroutines do I need for per disk? So need to implement a tool to measure
the disk and generate the configs.

## io_uring

Important feature will be added in Linux Kernel 5.1.

## Data Cache

1. Cache scope (ZBuf/ZCold/all)

2. LRU MRU (Multi or smart policy)

3. Prefetch
