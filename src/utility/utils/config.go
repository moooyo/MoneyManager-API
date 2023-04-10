package utils

import (
	"MoneyManager/src/model"
	"flag"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var configInstance *model.Config
var configOnce sync.Once

func init() {
	// Get config file path from flag
	configPath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	configOnce.Do(func() {
		config, err := LoadConfig(*configPath)
		if err != nil {
			// log.TrackError(nil, err)
			panic(err)
		}
		configInstance = config
	})
}

func GetConfig() *model.Config {
	return configInstance
}

func LoadConfig(path string) (*model.Config, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		// log.TrackError(nil, err)
		return nil, err
	}
	defer file.Close()
	var data []byte
	_, err = file.Read(data)
	if err != nil {
		// log.TrackError(nil, err)
		return nil, err
	}

	config, err := ParseConfig(data)
	if err != nil {
		// log.TrackError(nil, err)
		return nil, err
	}

	return config, nil
}

func ParseConfig(data []byte) (*model.Config, error) {
	var config model.Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		// log.TrackError(nil, err)
		return nil, err
	}
	return &config, nil
}
