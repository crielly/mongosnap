package command

import (
	"github.com/mitchellh/cli"
	"fmt"
	"flag"
)

// Backup command performs a Mongo backup
type Backup struct {
	MongoDir string
	UI cli.Ui
}

// Run the backup
func (b *Backup) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("backup", flag.ContinueOnError)
	cmdFlags.Usage = func() { 
		b.UI.Output(b.Help()) 
	}

	cmdFlags.StringVar(&b.MongoDir, "mongo-dir", "/data/db", "The Mongo Data Directory to backup")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	b.UI.Output(fmt.Sprintf("Run a backup on dir: %v", b.MongoDir))
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
