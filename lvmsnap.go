package main

import (
	"fmt"
	"log"
	"os/exec"

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
	arguments, _ := docopt.Parse(usage, nil, true, "mongosnap", false)
	size := arguments["<size>"].(string)
	name := arguments["<name>"].(string)
	path := arguments["<path>"].(string)
	fmt.Println(arguments)
	lvmSnap(size, name, path)
}

func lvmSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -L%sM -s -n %s %s", size, name, path)
	run := exec.Command(cmd)
	stdoutStderr, err := run.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
