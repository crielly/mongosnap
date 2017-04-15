package readconfig

import (
	"io/ioutil"

	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
)

// ReplicaConfig describes the configuration of a MongoD process
type ReplicaConfig struct {
	Replica string            `json:"replica"`
	Conf    map[string]string `json:"conf"`
}

// ReadConfig prints a test MongoD Replica Config
func ReadConfig(configPath string) (replconf ReplicaConfig, err error) {

	y, err := ioutil.ReadFile(configPath)
	logger.Error.Println(err)

	err = yaml.Unmarshal(y, &replconf)
	logger.Error.Println(err)

	return
}
