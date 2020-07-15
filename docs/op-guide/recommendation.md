Hardware & Software Requirements
===

## Platform

x86-64 architecture with AVX2 or newer

For:

1. Zai's Erasure Codes Lib need AVX2/AVX512 to gain high performance
2. Some basic libraries are using SIMD to accelerate .

## OS

Linux-4.4+

For:

1. Fix bug which causes containers that use veth devices for network routing to not check TCP checksum (v4.4-rc8)

    It's not a bad choice run Zai(it's a stateless entrance) on Docker, so this patch is important.
    
    Details are [here](https://tech.vijayp.ca/linux-kernel-bug-delivers-corrupt-tcp-ip-data-to-mesos-kubernetes-docker-containers-4986f88f7a19)
    
2. Only enforce stable page writes if the backing device requires it (v3.9-rc1)

    It's important to improve write performance.

    Details are [here](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=1d1d1a767206fbe5d4c69493b7e6d2a8d08cc0a0) 

ps:

Zai can run on Mac too.

## Filesystem

XFS is recommended.

Some good properties of XFS:

1. [Self-Describing Metadata](https://www.kernel.org/doc/Documentation/filesystems/xfs-self-describing-metadata.txt)

2. Neighbor Page Flushing, this feature will flush more dirty pages if we call sync_file_range helping
to improve perf because the unit of I/O is getting bigger.

## Clock

Zai relies on correct clock. NTP is needed.

## Server Recommendations

### Production environment 

| Component | CPU | Memory | Hard Disk | Network | Instance Number (Minimum Requirment) | Optional |
| :-----: | :------: | :------: |:-----: | :------: | :------: |:-----: |
| Zai | 8 core+ | 16 GB+ | SAS |  10 Gigabit network card (2 preferred) | 3 | No |
| Keeper | 8 core+ | 16 GB+ | NVMe SSD |  10 Gigabit network card (2 preferred) | 3 | No |
| ZBuf | 16 core+ | 64 GB+ | NVMe SSD |  10 Gigabit network card (2 preferred) | 3 | No |
| ZCold | 8 core+ | 32 GB+ | SATA HDD |  10 Gigabit network card (2 preferred) | 13 | Yes |

> **Note:**
>
> - In practice, it's a good choice to deploy at least 3 Zai node. 
    Help to reduce impact when nodes down.
>
> - Two CPU core for one NVMe SSD is a balanced choice in most cases.
>
> - QLC SSD is a good choice for ZBuf.
