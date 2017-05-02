package command

import (
	"github.com/mitchellh/cli"
	"fmt"
	"flag"
	"github.com/crielly/mongosnap/backconfig"
	"github.com/crielly/mongosnap/logger"
)
	// "github.com/crielly/mongosnap/replconfig"
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

	cmdFlags.StringVar(&b.BackConfYamlPath, "confpath", "./backconfig/mongosnap.yml", "Path to YAML MongoSnap config")

	if err := cmdFlags.Parse(args); err != nil {
		logger.Error.Println(err)
	}

	backconf, err := backconfig.BackConfig(b.BackConfYamlPath)

	if err != nil {
		logger.Error.Println(err)
	}

	for i, v := range backconf.Cluster.ReplicaConfs {
		fmt.Printf("Index %d, value %s\n", i,v)
	}

	// lvm.TakeSnap(conf.ClusSize, b.Name, fmt.Sprintf("%s%s", b.Volgrp, b.Lvol))
	
	// lvm.MountLvmSnap(fmt.Sprintf("%s%s", b.Volgrp, b.Name), b.MountPath, b.Filesystem, b.Opts)

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
