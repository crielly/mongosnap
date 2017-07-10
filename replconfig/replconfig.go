package replconfig

import (
	"io/ioutil"

	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
)

// Config struct describes the values we consider relevant for parsing
// out of a yaml mongodb conf file
type Config struct {
	Net struct {
		Port   int    `json:"port"`
		BindIP string `json:"bindIp"`
	} `json:"net"`
	Storage struct {
		DbPath string `json:"dbPath"`
	} `json:"storage"`
	Replication struct {
		ReplSetName string `json:"replSetName"`
	} `json:"replication"`
}

// ReplConfig unmarshals the Yaml from a mongodb.conf file and makes it available
// for use in a Config struct
func ReplConfig(configPath string) (replconf Config, err error) {

	y, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error.Println(err)
	}

	err = yaml.Unmarshal(y, &replconf)
	if err != nil {
		logger.Error.Println(err)
	}

	return replconf, err
}
