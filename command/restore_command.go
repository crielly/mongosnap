package command

import (
	"github.com/mitchellh/cli"
	"fmt"
)

// Restore command performs a Mongo restore
type Restore struct {
	UI cli.Ui
}

// Run the Restore
func (r *Restore) Run(args []string) int {
	r.UI.Output(fmt.Sprintf("Run a restore here with args: %v\n", args))
	return 0
}

// Help string for the restore command
func (r *Restore) Help() string {
	return "Help string for the restore command"
}

// Synopsis for the restore command
func (r *Restore) Synopsis() string {
	return "Synopsis string for the Restore command"
}
