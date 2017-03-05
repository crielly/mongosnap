# mongosnap - a MongoDB Cluster backup tool

**Note**: This thing isn't even remotely ready for use, run it at your own risk! I'm just writing this to solve a problem for my org - if it helps somebody else out, so much the better. Contributions/expertise/constructive criticism welcome.

#####Purpose:

Provide a tool for backing up large, sharded, distributed MongoDB deployments that doesn't rely on `mongodump` and uses cloud storage (ie S3) as a backend.

#####Motivation:

I have a large Mongo cluster at work that takes ~16 hours to dump via highly parallelized MongoDumps and much, much longer to restore from those dumps. We have an existing python-based behemoth of a solution that leverages a pile of EBS volumes, but it's unnecessarily complex due to handling those disks and is somewhat prone to failure - not to mention the fact that EBS costs a lot.

#####Method:

`mongosnap` will utilize volume/filesystem snapshots - LVM at first, though I'd love to support ZFS and eventually XFS-type FS snapshots as well - and streaming of data directly into S3 - to achieve fast, efficient backups across multiple hosts in a Mongo deployment of any size.
