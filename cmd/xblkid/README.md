xblkid
===

xblkid is similar with `blkid` command but focus on PARTUUID & Serial_Number.

## Usage

Print all devices.

```shell
xblkid
```

e.g.

```shell
./xblkid


-----------
dev_name: nvme0n1, S/N: PHLJ038600KK8P0HGN
no partition for this device
-----------
dev_name: nvme1n1, S/N: PHHH922506LJ256B
part_name: nvme1n1p1, mnt_path: /boot/efi, uuid: 0f9e0466-9f72-4ffa-b375-2aa8fdaa5b85
part_name: nvme1n1p2, mnt_path: /, uuid: 18449fc6-5f0d-4ee8-a350-d038d56579b0
-----------
dev_name: sda, S/N: VCHB5SZP
part_name: sda1, mnt_path: /data2, uuid: 3a20a8ad-4803-4612-bd76-42e406b01b02
-----------
dev_name: sdb, S/N: WD-WCC7K4FAT9HN
part_name: sdb1, mnt_path: /log, uuid: f0809b5c-f800-4c08-901c-55d2f1e45bad
part_name: sdb2, mnt_path: /data, uuid: 8b623c72-0536-4add-90de-1a0ce8c3af0e
```

Print a certain device.

```shell
xblkid -dev=<dev_name>
```

e.g.

```shell
./xblkid -dev=sdb

-----------
dev_name: sdb, S/N: WD-WCC7K4FAT9HN
part_name: sdb1, mnt_path: /log, uuid: f0809b5c-f800-4c08-901c-55d2f1e45bad
part_name: sdb2, mnt_path: /data, uuid: 8b623c72-0536-4add-90de-1a0ce8c3af0e
```