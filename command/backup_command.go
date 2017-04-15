package command

import (
	"github.com/mitchellh/cli"
	"fmt"
	"flag"
	"github.com/crielly/mongosnap/lvm"
	"github.com/crielly/mongosnap/s3upload"
	"log"
)

// Backup command performs a Mongo backup
type Backup struct {
	Volgrp string
	Lvol string
	MountPath string
	Filesystem string
	Opts string
	Size string
	Name string
	S3bucket string
	S3object string
	UI cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() { 
		b.UI.Output(b.Help()) 
	}

	cmdFlags.StringVar(&b.Volgrp, "vg", "/dev/ephem_vg/", "Path to volume group")
	cmdFlags.StringVar(&b.Lvol, "lvol", "ephem_lv", "The Mongo Data Directory to backup")
	cmdFlags.StringVar(&b.MountPath, "mountpath", "/backup", "dir to mount snap to and backup from")
	cmdFlags.StringVar(&b.Filesystem, "fs", "xfs", "Filesystem the snapshotted lvol is using")
	cmdFlags.StringVar(&b.Opts, "opts", "", "mount opts")
	cmdFlags.StringVar(&b.Name, "name", "mongobackup", "Name to give the snapshot")
	cmdFlags.StringVar(&b.Size, "size", "1024", "Size of the snapshot")
	cmdFlags.StringVar(&b.S3bucket, "s3bucket", "gazaro-operations", "S3 bucket in which to store backup")
	cmdFlags.StringVar(&b.S3object, "s3object", "mongo-backups/mongosnap/backup", "Path within S3 Bucket")

	if err := cmdFlags.Parse(args); err != nil {
		log.Fatal(err)
	}

	lvm.TakeSnap(b.Size, b.Name, fmt.Sprintf("%s%s", b.Volgrp, b.Lvol))
	
	lvm.MountLvmSnap(fmt.Sprintf("%s%s", b.Volgrp, b.Name), b.MountPath, b.Filesystem, b.Opts)

	s3upload.Zip(b.MountPath, b.S3bucket, b.S3object)

	return 0
}

// Help string for the backup command
func (b *Backup) Help() string {
	return "Help string for the backup command"
}

// Synopsis string for the backup command
func (b *Backup) Synopsis() string {
	return "Synopsis string for the Backup command"
}
