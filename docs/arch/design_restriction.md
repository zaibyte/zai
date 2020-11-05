Design Restriction
====

***"A lot of great creativity comes from restrictions."***        ---- Giles Duley

# Background

Zai's major job is to provide distributed file system for Tesamc's machine learning platform.

Public could & local NAS (maybe depracated) will be the backup storage, so Zai won't implement warm/cold storage tier.

# Targets

>- Capacity: Each box 4PB at most
>
>   We could have 255 boxes in a single namespace, which means about 1 EB.
>
>- Blazing throughput for big files
>
>   Most of the files are big files(> 100 MB)
>
>- Low latency for small files
>
>   There are some requests are quite small(e.g. request for a group physical address), Zai should be blocked on these requests.
>

## Restrictions

### High Availability

1. ZBuf need replicas
2. Keeper need be based on consensus algorithm 
3. Entrance should be stateless

### Low Latency

1. RPC framework should detect request which need be low latency
2. I/O stack should be as short as possible:
    
    a. Combine file APIs and object APIs.
    
    b. Try to use main memory fully.
    
    c. Combine file meta and object meta.
    
3. Meta should be extremely fast, lock-free(wait-free is better) for reading operation.
    
### High Throughput

1. Split big files into pieces for parallel I/O
2. Automatically prefetch blocks for next reading

### Focus on hundreds TB and millions files

For 256 TB usable capacity, 32 MB each file, there will be about 10 millions files.
The total meta size should be < 1GB, which means 128B/file. Good enough for Tesamc's needs.


                            
                           