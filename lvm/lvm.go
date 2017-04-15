package lvm

import (
	"fmt"
	"os/exec"
	"os"
	"github.com/crielly/mongosnap/logger"
	"log"
)

// TakeSnap triggers an LVM snapshot
func TakeSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -L %sM -s -n %s %s", size, name, path)
	fmt.Println(cmd)
	run := exec.Command("bash", "-c", cmd)
	_, err := run.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}
}

func makeDir(path string, mode os.FileMode) {
	err := os.MkdirAll(path, mode)
	if err != nil {
		log.Fatal(err)
	} else {
		logger.Info.Printf("Created directory %s with permissions %s", path, mode)
	}
}

// MountLvmSnap will mount the Snapshot to filesystem
func MountLvmSnap(snappath, mountpath, fstype, opts string) {
	makeDir(mountpath, 0644)

	cmd := fmt.Sprintf("mount -t %s %s %s %s", fstype, snappath, mountpath, opts)

	logger.Info.Println(cmd)

	run := exec.Command("bash", "-c", cmd)

	_, err := run.CombinedOutput()

	if err != nil {
		logger.Error.Println(err)
	} else {
		logger.Info.Printf("Mounted snapshot %s at path %s using opt string %s", snappath, mountpath, opts)
	}
	
}
