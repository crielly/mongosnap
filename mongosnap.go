// Mongosnap is used for backing up a MongoDB cluster via volume snapshots.
// It supports LVM and ZFS snapshots to achieve a consistent state backup
// of a MongoDB cluster of any size or sharding configuration
// It is designed to then stream the snapshot directly into S3
// without ever writing the data to disk, ensuring that even a large
// and very full cluster can be backed up without using additional disk resources.

package main

import (
	"log"
	"github.com/mitchellh/cli"
	"os"
	"github.com/crielly/mongosnap/command"
	"fmt"
)
func main() {
	// Call realMain instead of doing the work here so we can use
	// `defer` statements within the function and have them work properly.
	// (defers aren't called with os.Exit)
	os.Exit(realMain())
}

func realMain() int {

	ui := &cli.BasicUi{
		Reader:			os.Stdin,
		Writer:			os.Stdout,
		ErrorWriter: 	os.Stderr,
	}

	c := cli.NewCLI("mongosnap", "0.0.1")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"backup": func() (cli.Command, error) {
			return &command.Backup{
				UI: &cli.ColoredUi{
					Ui:	ui,
					OutputColor: cli.UiColorBlue,
				},
			}, nil
		},
		"restore": func() (cli.Command, error) {
			return &command.Restore{
				UI: &cli.ColoredUi{
					Ui:	ui,
					OutputColor: cli.UiColorYellow,
				},
			}, nil
		},
	}
	fmt.Println("Print some Args: ", c.Args)

	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
		return 1
	}

	return exitStatus

}

