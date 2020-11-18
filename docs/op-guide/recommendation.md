Hardware & Software Requirements
===

## Platform

x86-64 architecture with AVX2 or newer

For:

1. Zai's Erasure Codes Lib need AVX2/AVX512 to gain high performance
2. Some basic libraries are using SIMD to accelerate .

## OS

CentOS8 & Linux-4.19 (LTS) 

**Only choose the LTS version of kernel.**

For:

1. Fix bug which causes containers that use veth devices for network routing to not check TCP checksum (v4.4-rc8)

    It's not a bad choice run Zai(it's a stateless entrance) on Docker, so this patch is important.
    
    Details are [here](https://tech.vijayp.ca/linux-kernel-bug-delivers-corrupt-tcp-ip-data-to-mesos-kubernetes-docker-containers-4986f88f7a19)
    
2. Only enforce stable page writes if the backing device requires it (v3.9-rc1)

    It's important to improve write performance.

    Details are [here](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=1d1d1a767206fbe5d4c69493b7e6d2a8d08cc0a0) 

ps:

Zai could run on Mac too.

## IO Uring

In the next, we may use IO Uring to accelerate the I/O. If we want this, Linux-5.4+(LTS) is a better choice.

## FUSE

MAX_PAGES option supports was added in Linux-4.20, help to improve performance

## Filesystem

XFS is recommended.

Some good properties of XFS:

1. [Self-Describing Metadata](https://www.kernel.org/doc/Documentation/filesystems/xfs-self-describing-metadata.txt)

2. Neighbor Page Flushing, this feature will flush more dirty pages if we call sync_file_range helping
to improve perf because the unit of I/O is getting bigger.

## Clock

Zai relies on correct clocks. NTP is needed.

## Server Recommendations

### Production environment 

| Component | CPU | Memory | Hard Disk | Network | Instance Number (Minimum Requirment) | Optional |
| :-----: | :------: | :------: |:-----: | :------: | :------: |:-----: |
| Zai | 8 core+ | 16 GB+ | SAS |  10 Gigabit network card (2 preferred) | 3 | No |
| Keeper | 8 core+ | 16 GB+ | NVMe SSD |  10 Gigabit network card (2 preferred) | 3 | No |
| ZBuf | 16 core+ | 64 GB+ | NVMe SSD |  10 Gigabit network card (2 preferred) | 3 | No |

> **Note:**
>
> - In practice, it's a good choice to deploy at least 3 Zai node. 
    Help to reduce impact when nodes down.
>
> - 500MB_IO/s/core is a better start in most cases.
>
> - TLC/QLC SSD is a good choice for ZBuf.