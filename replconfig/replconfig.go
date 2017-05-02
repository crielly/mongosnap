package replconfig

import (
	"io/ioutil"

	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
)

// Config describes the configuration of a MongoD process
type Config struct {
	Net struct {
		Port int `json:"port"`
		BindIP string `json:"bindIp"`
	} `json:"net"`
	Storage struct {
		DbPath string `json:"dbPath"`
	} `json:"storage"`
}

// ReplConfig prints a test MongoD Replica Config
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
