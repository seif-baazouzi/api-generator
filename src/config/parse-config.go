package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func Parse(configPath string) (Config, error) {
	var config Config

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
