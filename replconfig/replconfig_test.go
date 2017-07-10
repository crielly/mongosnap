package replconfig

import (
	"testing"
)

// We want to read in the contents of mongodb.conf in this directory using
// the ReplConfig function and then test the values read in against known values
// defined here in the test.
func TestReplNetConfig(t *testing.T) {
	replconf, err := ReplConfig("mongodb.conf")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Replica Config: %v", err)
	}

	if replconf.Net.Port != 27017 {
		t.Errorf("Incorrect Port - expected 27017, got %d", replconf.Net.Port)
	}

}

func TestReplStorageConfig(t *testing.T) {
	replconf, err := ReplConfig("mongodb.conf")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Replica Config: %v", err)
	}

	if replconf.Storage.DbPath != "/data/db00" {
		t.Errorf("Incorrect DBPath - expected /data/db00, got %s", replconf.Storage.DbPath)
	}

}

func TestReplReplicationConfig(t *testing.T) {
	replconf, err := ReplConfig("mongodb.conf")

	if err != nil {
		t.Fatalf("Encountered an error parsing the Replica Config: %v", err)
	}

	if replconf.Replication.ReplSetName != "prodreplica08" {
		t.Errorf("Incorrect ReplSetName - expected prodreplica08, got %s", replconf.Replication.ReplSetName)
	}

}