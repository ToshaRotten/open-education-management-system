package database_configurator

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Host int
}

//Parse file ...
func ParseConfigFile(path string) (*Config, error) {
	var c Config
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
