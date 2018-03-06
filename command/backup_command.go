package command

import (
	"flag"
	"fmt"
	"github.com/crielly/mongosnap/backupconfig"
	"github.com/crielly/mongosnap/logger"
	"github.com/crielly/mongosnap/lvm"
	"github.com/crielly/mongosnap/replconfig"
	"github.com/crielly/mongosnap/s3upload"
	"github.com/mitchellh/cli"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Backup command performs a Mongo backup
type Backup struct {
	BackConfYamlPath string
	SnapConfYamlPath string
	UI               cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		b.UI.Output(b.Help())
	}

	cmdFlags.StringVar(
		&b.BackConfYamlPath,
		"confpath",
		"backupconfig/backupconfig.yml",
		"Path to YAML MongoSnap config",
	)

	cmdFlags.StringVar(
		&b.SnapConfYamlPath,
		"confpath",
		"lvm/snapconfig.yml",
		"Path to YAML LVM Snapshot Configuration file",
	)

	if err := cmdFlags.Parse(args); err != nil {
		logger.Error.Println(err)
	}

	backconf, err := backupconfig.BackupConfig(b.BackConfYamlPath)
	if err != nil {
		logger.Error.Printf(
			"Error reading backup config: %s",
			err,
		)
		os.Exit(1)
	}

	snapconf, err := lvm.SnapConfig(b.SnapConfYamlPath)
	if err != nil {
		logger.Error.Printf(
			"Error reading snapshot config: %s",
			err,
		)
		os.Exit(1)
	}

	volpath := filepath.Join(
		backconf.Cluster.Storage.VolumeGroup,
		backconf.Cluster.Storage.LogicalVolume,
	)

	snappath := filepath.Join(
		snapconf.Snapshot.VgName,
		snapconf.Snapshot.SnapshotName,
	)

	logger.Info.Printf("Logical Volume Path: %s", volpath)
	logger.Info.Printf("LVM Snapshot path: %s", snappath)

	lvm.MakeDir(snapconf)

	lvm.TakeSnap(snapconf)

	lvm.MountLvmSnap(snapconf)

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

			dbpath := filepath.Join(
				snapconf.Snapshot.MountPath,
				replconf.Storage.DbPath,
			)

			s3obj := filepath.Join(
				backconf.S3.ObjectPath,
				replconf.Replication.ReplSetName,
				time.Now().Format("20060102150405"),
			)

			s3upload.Zip(dbpath, backconf.S3.Bucket, s3obj)

		}(v)
	}
	wg.Wait()

	lvm.Cleanup(snapconf)

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
