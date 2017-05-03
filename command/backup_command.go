package command

import (
	"github.com/mitchellh/cli"
	"fmt"
	"flag"
	"github.com/crielly/mongosnap/backconfig"
	"github.com/crielly/mongosnap/logger"
	"sync"
	"github.com/crielly/mongosnap/replconfig"
)
	// "github.com/crielly/mongosnap/lvm"
	// "github.com/crielly/mongosnap/s3upload"

// Backup command performs a Mongo backup
type Backup struct {
	BackConfYamlPath string
	UI cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() { 
		b.UI.Output(b.Help()) 
	}

	cmdFlags.StringVar(&b.BackConfYamlPath, "confpath", "backconfig/mongosnap.yml", "Path to YAML MongoSnap config")

	if err := cmdFlags.Parse(args); err != nil {
		logger.Error.Println(err)
	}

	backconf, err := backconfig.BackConfig(b.BackConfYamlPath)

	if err != nil {
		logger.Error.Println(err)
	}

	snapsize := backconf.Cluster.Snapshot.Size
	snapname := backconf.Cluster.Snapshot.SnapshotName
	volpath := fmt.Sprintf("%s/%s", backconf.Cluster.Storage.VolumeGroup, backconf.Cluster.Storage.LogicalVolume)
	snappath := fmt.Sprintf("%s/%s", volpath, snapname)
	mountpath := backconf.Cluster.Snapshot.MountPath
	s3bucket := backconf.S3.Bucket
	s3Object := backconf.S3.ObjectPath
	fstype := backconf.Cluster.Storage.FileSystem
	mountopts := backconf.Cluster.Snapshot.Opts

	fmt.Println("S3 Bucket: ", s3bucket)
	fmt.Println("S3 Path: ", s3Object)
	fmt.Println("Snapshot Name: ", snapname)
	fmt.Println("Snapshot Size: ", snapsize)
	fmt.Println("Volume Path to Snapshot: ", volpath)
	fmt.Println("Snapshot Path: ", snappath)
	fmt.Println("Snapshot Mount Path: ", mountpath)
	fmt.Println("Filesystem type: ", fstype)
	fmt.Println("Mount Options: ", mountopts)
	// lvm.Cleanup(snappath, mountpath)

	// lvm.TakeSnap(snapsize, snapname, snappath)

	// lvm.MountLvmSnap(snappath, mountpath, fstype, mountopts)

	var wg sync.WaitGroup
	wg.Add(len(backconf.Cluster.ReplicaConfs))

	for _, v := range backconf.Cluster.ReplicaConfs {
		go func(v string) {
			defer wg.Done()

			replconf, err := replconfig.ReplConfig(v)

			if err != nil {
				logger.Error.Println(err)
			}

			dbpath := fmt.Sprintf("%s%s", mountpath, replconf.Storage.DbPath)
			fmt.Println(dbpath)
			
			// s3upload.Zip(dbpath, s3bucket, s3Object)

		}(v)
	}

	wg.Wait()
	// lvm.Cleanup(snappath, mountpath)

	// s3upload.Zip(b.MountPath, b.S3bucket, b.S3object)

	// lvm.LvmCleanup(fmt.Sprintf("%s%s", b.Volgrp, b.Name), b.MountPath)

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
