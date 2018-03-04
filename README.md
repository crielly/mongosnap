# mongosnap

## A MongoDB Cluster backup tool

**Note**: This thing isn't even remotely ready for use, run it at your own risk! I'm just writing this to solve a problem for my org - if it helps somebody else out, so much the better. Contributions/expertise/constructive criticism welcome (you can use https://gitter.im/mongosnap/Lobby)

### Purpose

Provide a tool for backing up large, sharded, distributed MongoDB deployments that doesn't rely on `mongodump` and uses cloud storage (ie S3) as a backend.

### Method

`mongosnap` will utilize volume/filesystem snapshots - LVM at first, though I'd love to support ZFS and eventually XFS-type FS snapshots as well - and streaming of data directly into S3 - to achieve fast, efficient backups across multiple hosts in a Mongo deployment of any size.

### Additional Information

Consult the [Github Wiki](https://github.com/crielly/mongosnap/wiki)