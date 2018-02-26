package backupconfig

import (
	"testing"
)

func TestBackupStorageConfig(t *testing.T) {
	backconf, err := BackupConfig("backupconfig.yml")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Backup config: %v", err)
	}

	if backconf.Cluster.Storage.VolumeGroup != "/dev/ephem_vg" {
		t.Errorf("Incorrect Volume Group - expected /dev/ephem_vg, got %s", backconf.Cluster.Storage.VolumeGroup)
	}

	if backconf.Cluster.Storage.LogicalVolume != "ephem_lv" {
		t.Errorf("Incorrect Logical Volume - expected ephem_lv, got %s", backconf.Cluster.Storage.LogicalVolume)
	}

	if backconf.Cluster.Storage.FileSystem != "xfs" {
		t.Errorf("Incorrect Filesystem - expected xfs, got %s", backconf.Cluster.Storage.FileSystem)
	}

}

func TestBackupSnapConfig(t *testing.T) {
	backconf, err := BackupConfig("backupconfig.yml")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Backup config: %v", err)
	}

	if backconf.Cluster.Snapshot.MountPath != "/backup" {
		t.Errorf("Incorrect Mount Path - expected /backup, got %s", backconf.Cluster.Snapshot.MountPath)
	}

	if backconf.Cluster.Snapshot.SnapshotName != "mongobackup" {
		t.Errorf("Incorrect snapshot name - expected mongobackup, got %s", backconf.Cluster.Snapshot.SnapshotName)
	}

	if backconf.Cluster.Snapshot.Size != "100%FREE" {
		t.Errorf("Incorrect snapshot size - expected 100%%FREE, got %s", backconf.Cluster.Snapshot.Size)
	}

	if backconf.Cluster.Snapshot.Opts != "-onouuid,ro" {
		t.Errorf("Incorrect snapshot size - expected -onouuid,ro but got %s", backconf.Cluster.Snapshot.Opts)
	}

}

func TestBackupS3Config(t *testing.T) {
	backconf, err := BackupConfig("backupconfig.yml")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Backup config: %v", err)
	}

	if backconf.S3.Bucket != "mongo-db-backups" {
		t.Errorf("Incorrect S3 Bucket name - expected mongo-db-backups, got %s", backconf.S3.Bucket)
	}

	if backconf.S3.ObjectPath != "mongo-db-backups-test" {
		t.Errorf("Incorrect S3 object path - expected mongo-db-backups-test, got %s", backconf.S3.ObjectPath)
	}
}

func TestBackupReplConfs(t *testing.T) {
	backconf, err := BackupConfig("backupconfig.yml")

	confs := []string{"/etc/mongodb00.conf", "/etc/mongodb01.conf", "/etc/mongodb02.conf"}

	if err != nil {
		t.Fatalf("Encountered an error parsing the Backup config: %v", err)
	}

	check := func(a string, list []string) bool {
		for _, b := range list {
			if b == a {
				return true
			}
		}
		return false
	}

	for _, a := range confs {
		exists := check(a, backconf.Cluster.ReplicaConfs)
		if exists != true {
			t.Errorf("Expected value not found in ReplicaConfigs list: %s", a)
		}
	}
}
