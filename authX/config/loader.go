package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var loadedConfigs *Configs

func LoadAll() (Configs, error) {
	configs := Configs{}

	// loading config.yml
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return configs, errors.New("Failed to read config file! reason: " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		return configs, errors.New("Failed to unmarshal yml file! reason: " + err.Error())
	}

	loadedConfigs = &configs
	return configs, nil
}

func GetConfigs() (Configs, error) {
	if loadedConfigs == nil {
		LoadAll()
	}
	if loadedConfigs == nil {
		return Configs{}, errors.New("no config loaded")
	}
	conf := *loadedConfigs
	return conf, nil
}
