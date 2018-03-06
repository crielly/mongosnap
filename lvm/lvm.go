package lvm

import (
	"fmt"
	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
	lvm "github.com/nak3/go-lvm"
	"io/ioutil"
	"os"
	"syscall"
)

// Config is the struct which defines the data expected from our config yaml
type Config struct {
	Snapshot struct {
		MountPath    string      `json:"mountPath"`
		MountPerms   os.FileMode `json:"mountPerms"`
		MountFlags   string      `json:"mountflags"`
		MountData    string      `json:"mountdata"`
		SnapshotName string      `json:"snapshotName"`
		Size         uint64      `json:"size"`
		VgName       string      `json:"vgname"`
		LvName       string      `json:"lvname"`
		FSType       string      `json:"fstype"`
	} `json:"snapshot"`
}

// SnapConfig gets snapshot configuration from yaml
func SnapConfig(configpath string) (snapconf Config, err error) {
	s, err := ioutil.ReadFile(configpath)
	if err != nil {
		logger.Error.Println(err)
	}

	err = yaml.Unmarshal(s, &snapconf)
	if err != nil {
		logger.Error.Println(err)
	}

	return snapconf, err
}

// FindLV finds the logical volume matching the name and Volume Group supplied in config
func FindLV(vgname, lvname string) (lvo *lvm.LvObject, vgo *lvm.VgObject, err error) {
	vglist := lvm.ListVgNames()

	for _, v := range vglist {
		vgo := &lvm.VgObject{}
		vgo.Vgt = lvm.VgOpen(v, "w")

		if vgo.GetName() == vgname {
			logger.Info.Printf(
				"Found VG matching name %s: %+v",
				vgname,
				vgo.GetUuid(),
			)

			lvo, err = vgo.LvFromName(lvname)

			if err != nil {
				logger.Error.Printf(
					"Error finding Logical Volume by Name: %s",
					err,
				)
			} else {
				logger.Info.Printf(
					"Found LV matching name %s: %+v",
					lvname,
					lvo.GetUuid(),
				)
				return lvo, vgo, err
			}

		}
		logger.Info.Printf("Just checking if we make it this far")
	}
	panic("Couldn't locate the specified Logical Volume")
}

// TakeSnap triggers an LVM snapshot
func TakeSnap(snapconf Config) (snapshot *lvm.LvObject) {

	// Get an LvObject
	lvo, vgo, err := FindLV(
		snapconf.Snapshot.VgName,
		snapconf.Snapshot.LvName,
	)

	if err != nil {
		logger.Error.Printf(
			"Error finding LV %s in VG %s: %s",
			snapconf.Snapshot.VgName,
			snapconf.Snapshot.LvName,
			err,
		)
		os.Exit(1)
	}

	logger.Info.Printf(
		"Using LV %s for snapshot %s",
		lvo.GetUuid(),
		snapconf.Snapshot.SnapshotName,
	)

	freesize := uint64(vgo.GetFreeSize())

	snapobject, err := lvo.Snapshot(
		snapconf.Snapshot.SnapshotName,
		freesize,
	)

	gp, _ := snapobject.GetProperty("lv_path")

	if err != nil {
		logger.Error.Printf("Error Creating Snapshot: %s", err)
		os.Exit(1)
	}

	logger.Info.Printf(
		"Snapshot created successfully with UUID %s - %+v",
		snapobject.GetUuid(),
		&gp,
	)

	return snapobject
}

// MakeDir creates the directory specified in config for mounting the snapshot
func MakeDir(snapconf Config) {
	err := os.MkdirAll(
		snapconf.Snapshot.MountPath,
		snapconf.Snapshot.MountPerms,
	)
	if err != nil {
		logger.Error.Printf(
			"Error creating directory %s: %s, exiting",
			snapconf.Snapshot.MountPath,
			err,
		)
		os.Exit(1)
	} else {
		logger.Info.Printf(
			"Created directory %s with permissions %s",
			snapconf.Snapshot.MountPath,
			snapconf.Snapshot.MountPerms,
		)
	}
}

// MountLvmSnap will mount the Snapshot to filesystem
func MountLvmSnap(snapconf Config) {

	lvpath := fmt.Sprintf(
		"/dev/%s/%s",
		snapconf.Snapshot.VgName,
		snapconf.Snapshot.SnapshotName,
	)

	err := syscall.Mount(
		lvpath,
		snapconf.Snapshot.MountPath,
		snapconf.Snapshot.FSType,
		syscall.MS_RDONLY,
		snapconf.Snapshot.MountData,
	)

	if err != nil {
		logger.Error.Printf("Error mounting snapshot: %s", err)
		os.Exit(1)
	} else {
		logger.Info.Printf(
			"Successfully mounted %s at %s",
			lvpath,
			snapconf.Snapshot.MountPath,
		)
	}
}

// Cleanup unmounts and deletes our snapshot and deletes the directory we used to mount it
func Cleanup(snapconf Config) {

	logger.Info.Printf("Running Cleanup")

	err := syscall.Unmount(
		snapconf.Snapshot.MountPath,
		syscall.MNT_FORCE,
	)

	if err != nil {
		logger.Info.Printf(
			"Attempted unmounting snapshot from mountpath: %s",
			err,
		)
	} else {
		logger.Info.Printf(
			"Successfully unmounted snapshot from %s",
			snapconf.Snapshot.MountPath,
		)
	}

	snapobject, _, err := FindLV(
		snapconf.Snapshot.VgName,
		snapconf.Snapshot.SnapshotName,
	)

	if err != nil {
		logger.Info.Printf(
			"Couldn't find relevant snapshot for deletion: %s",
			err,
		)
	} else {
		logger.Info.Printf(
			"Found LV %s during cleanup, will delete",
			snapobject.GetUuid(),
		)
	}

	err = snapobject.Remove()

	if err != nil {
		logger.Info.Printf(
			"Error removing snapshot: %s",
			err,
		)
	} else {
		logger.Info.Printf(
			"Successfully removed snapshot",
		)
	}

	err = os.RemoveAll(snapconf.Snapshot.MountPath)
	if err != nil {
		logger.Info.Printf(
			"Error removing MountPath %s: %s",
			snapconf.Snapshot.MountPath,
			err,
		)
	} else {
		logger.Info.Printf(
			"Successfully removed directory %s",
			snapconf.Snapshot.MountPath,
		)
	}
}
