package main

import (
	"fmt"

	"github.com/crielly/mongosnap/s3upload"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `mongosnap

Usage:
	mongosnap --snapname=(<snapname>) --size=(<size>) --path=(<path>)

Options:
	-h --help	Show usage information
	-n --snapname	Snapshot name
	-s --size	Size of the snapshot
	-p --snappath	Path to snap
`
	arguments, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		fmt.Println(err)
	}

	size := arguments["--size"].(string)
	name := arguments["--name"].(string)
	path := arguments["--path"].(string)

	fmt.Println(size, name, path)

	// lvmsnap.LvmSnap(size, name, path)

	s3upload.S3upload()

}
