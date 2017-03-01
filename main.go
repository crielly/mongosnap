package main

import (
	"fmt"

	"github.com/crielly/mongosnap/lvmsnap"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `mongosnap

Usage:
	mongosnap --name=<name> --size=<size> --path=<path>

Options:
	-h --help	Show usage information
	-n --name	Snapshot name
	-s --size	Size of the snapshot
	-p --path	Path to snap
`
	arguments, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		fmt.Println(err)
	}

	size := arguments["--size"].(string)
	name := arguments["--name"].(string)
	path := arguments["--path"].(string)

	lvmsnap.LvmSnap(size, name, path)

}
