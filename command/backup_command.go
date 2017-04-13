package command

import (
	"github.com/mitchellh/cli"
	"fmt"
	"flag"
	"github.com/crielly/mongosnap/lvm"
	"log"
)

// Backup command performs a Mongo backup
type Backup struct {
	MongoDir string
	MountPath string
	Filesystem string
	opts string
	Size string
	Name string
	UI cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() { 
		b.UI.Output(b.Help()) 
	}

	cmdFlags.StringVar(&b.MongoDir, "mongo-dir", "/data/db", "The Mongo Data Directory to backup")
	cmdFlags.StringVar(&b.MountPath, "mountpath", "/backup", "dir to mount snap to and backup from")
	cmdFlags.StringVar(&b.Filesystem, "filesystem", "xfs", "Filesystem the snapshotted lvol is using")
	cmdFlags.StringVar(&b.opts, "opts", "-onouuid,ro", "mount opts")
	cmdFlags.StringVar(&b.Name, "name", "snapshot", "Name to give the snapshot")
	cmdFlags.StringVar(&b.Size, "size", "1024", "Size of the snapshot")

	if err := cmdFlags.Parse(args); err != nil {
		log.Fatal(err)
	}

	b.UI.Output(fmt.Sprintf("Run a backup on dir: %v", b.MongoDir))
	lvm.TakeSnap(b.Size, b.Name, b.MongoDir)
	lvm.MountLvmSnap(b.MongoDir, b.MountPath, b.Filesystem, b.opts)

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
