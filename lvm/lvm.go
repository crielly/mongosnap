package lvm

import (
	"fmt"
	"os/exec"
	"os"
	"syscall"
	"log"
)

// TakeSnap triggers an LVM snapshot
func TakeSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -L %sM -s -n %s %s", size, name, path)
	fmt.Println(cmd)
	run := exec.Command("bash", "-c", cmd)
	stdoutStderr, err := run.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)

	if err != nil {
		log.Fatal(err)
	}
}

func makeDir(path string, mode os.FileMode) {
	err := os.MkdirAll(path, mode)
	if err != nil {
		log.Fatal(err)
	}
}

// MountLvmSnap will mount the Snapshot to filesystem
func MountLvmSnap(snappath, mountpath, filesystem, opts  string) {
	makeDir(mountpath, 0644)

	if err := syscall.Mount(snappath, mountpath, filesystem, 0, opts); err != nil {
		log.Fatal(err)
	}
	
}
