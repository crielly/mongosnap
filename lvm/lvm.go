package lvm

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/crielly/mongosnap/logger"
)

// TakeSnap triggers an LVM snapshot
func TakeSnap(size, name, path string) {
	cmd := fmt.Sprintf("lvcreate -l%s -s -n %s %s", size, name, path)
	fmt.Println(cmd)
	run := exec.Command("bash", "-c", cmd)
	output, err := run.CombinedOutput()

	logger.Info.Println(output)

	if err != nil {
		logger.Error.Println(err)
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

	cmd := fmt.Sprintf("mount %s %s %s", snappath, mountpath, opts)

	logger.Info.Println(cmd)

	run := exec.Command("bash", "-c", cmd)

	_, err := run.CombinedOutput()

	if err != nil {
		logger.Error.Println(err)
		Cleanup(snappath, mountpath)
	} else {
		logger.Info.Printf("Mounted snapshot %s at path %s using opt string %s", snappath, mountpath, opts)
	}

}

// Cleanup any snapshot at specified dir
func Cleanup(snappath, mountdir string) {

	err := syscall.Unmount(mountdir, 0)
	if err != nil {
		logger.Error.Println(err)
	}

	removesnapcmd := fmt.Sprintf("lvremove -f %s", snappath)
	logger.Info.Println(removesnapcmd)
	run := exec.Command("bash", "-c", removesnapcmd)
	_, err = run.CombinedOutput()
	if err != nil {
		logger.Error.Println(err)
	}

	err = os.RemoveAll(mountdir)
	if err != nil {
		logger.Error.Println(err)
	}
}
