package grpc_server

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// ParseConfigFile Parse configuration file ...
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
