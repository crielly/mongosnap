package readconfig

import (
	"fmt"

	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
)

// ReplicaConfig describes the configuration of a MongoD process
type ReplicaConfig struct {
	Name      string `json:"name"`
	Port      int    `json:"port"`
	Directory string `json:"directory"`
}

// ReadConfig prints a test MongoD Replica Config
func ReadConfig() {
	b2b := ReplicaConfig{"b2bReplica", 27017, "/data/db00"}
	y, err := yaml.Marshal(b2b)
	if err != nil {
		logger.LogError(err)

		fmt.Println(string(y))

	}
}
