package main

import (
	"fmt"
	"time"

	"github.com/crielly/mongosnap/logger"
	"github.com/crielly/mongosnap/readconfig"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `mongosnap

Usage:
	mongosnap --snapname=<snapname> --snapsize=<snapsize> --snappath=<snappath> --filepath=<filepath> --bucket=<bucket> --object=<object> --configpath=<configpath>

Options:
	-h --help		Show usage information
	-n --snapname	Snapshot name
	-s --snapsize	Size of the snapshot
	-p --snappath	Path to snap
	-f --filepath	File to archive
	-b --bucket		S3 bucket
	-o --object		S3 object path
	-c --configpath Path to yaml config
`
	arguments, err := docopt.Parse(usage, nil, true, "", false)
	logger.LogError(err)

	snapsize := arguments["--snapsize"].(string)
	snapname := arguments["--snapname"].(string)
	snappath := arguments["--snappath"].(string)
	filepath := arguments["--filepath"].(string)
	bucket := arguments["--bucket"].(string)
	object := arguments["--object"].(string)
	configpath := arguments["--configpath"].(string)

	t := time.Now().UTC().Format("20060102150405")
	fmt.Println(t)
	fmt.Println(snapsize, snapname, snappath)
	fmt.Println(filepath, bucket, object)

	// lvmsnap.LvmSnap(size, name, path)

	// s3upload.Zip(filepath, bucket, object)
	yamlconf, err := readconfig.ReadConfig(configpath)

	fmt.Println(yamlconf)

}
