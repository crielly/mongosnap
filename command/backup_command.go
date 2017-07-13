package command

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"os"
	"github.com/crielly/mongosnap/backupconfig"
	"github.com/crielly/mongosnap/logger"
	"github.com/crielly/mongosnap/lvm"
	"github.com/crielly/mongosnap/replconfig"
	"github.com/crielly/mongosnap/s3upload"
	"github.com/mitchellh/cli"
)

// Backup command performs a Mongo backup
type Backup struct {
	BackConfYamlPath string
	UI               cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		b.UI.Output(b.Help())
	}

	cmdFlags.StringVar(&b.BackConfYamlPath,
		"confpath",
		"backupconfig/backupconfig.yml",
		"Path to YAML MongoSnap config",
	)

	if err := cmdFlags.Parse(args); err != nil {
		logger.Error.Println(err)
	}

	backconf, err := backupconfig.BackupConfig(b.BackConfYamlPath)

	if err != nil {
		logger.Error.Println(err)
		os.Exit(1)
	}

	volpath := fmt.Sprintf("%s/%s",
		backconf.Cluster.Storage.VolumeGroup,
		backconf.Cluster.Storage.LogicalVolume,
	)

	snappath := fmt.Sprintf("%s/%s",
		backconf.Cluster.Storage.VolumeGroup,
		backconf.Cluster.Snapshot.SnapshotName,
	)

	logger.Info.Printf("Logical Volume Path: %s", volpath)
	logger.Info.Printf("LVM Snapshot path: %s", snappath)

	lvm.Cleanup(snappath,
		backconf.Cluster.Snapshot.MountPath,
	)

	lvm.TakeSnap(backconf.Cluster.Snapshot.Size,
		backconf.Cluster.Snapshot.SnapshotName,
		volpath,
	)

	lvm.MountLvmSnap(snappath,
		backconf.Cluster.Snapshot.MountPath,
		backconf.Cluster.Storage.FileSystem,
		backconf.Cluster.Snapshot.Opts,
	)

	var wg sync.WaitGroup
	wg.Add(len(backconf.Cluster.ReplicaConfs))

	for _, v := range backconf.Cluster.ReplicaConfs {
		go func(v string) {
			defer wg.Done()

			replconf, err := replconfig.ReplConfig(v)

			fmt.Println(replconf)

			if err != nil {
				logger.Error.Println(err)
			}

			dbpath := fmt.Sprintf("%s%s",
				backconf.Cluster.Snapshot.MountPath,
				replconf.Storage.DbPath,
			)

			s3obj := fmt.Sprintf("%s/%s/%s",
				backconf.S3.ObjectPath,
				replconf.Replication.ReplSetName,
				time.Now().Format("20060102150405",
			))

			s3upload.Zip(dbpath, backconf.S3.Bucket, s3obj)

		}(v)
	}
	wg.Wait()

	lvm.Cleanup(snappath, backconf.Cluster.Snapshot.MountPath)

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
