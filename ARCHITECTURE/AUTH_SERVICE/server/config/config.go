package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func New() *Config {
	return &Config{
		Port: ":5671",
		Host: "localhost",
	}
}

// ParseFile ..
func (c *Config) ParseFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}
