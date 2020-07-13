Durability
===

## Memory

ECC memory is needed.

## Disk

Please use enterprise devices. It's good for both durability and performance.

There are many reasons to make that choice, and I will supplement some which people always ignore:

As fas as I know, XFS(which is used in Zai) enable write barrier by default, that means when we call fsync, it not just flush the dirty pages, it also trigger block device's flush cmd.
For most of enterprise SSD with data loss protection, fsync will return ok when the data is reach the buffer, so it's faster.

And these enterprise SSD also has data integrity protection.

## Switch

We should be aware of switch's Ethernet CRC errors, it may cause disaster.

## Replace aged devices

Old devices never die, they just fade away.