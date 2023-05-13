package buffer

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Buffer struct {
	Configs []Config `json:"configs" yaml:"configs"`
}

func New() *Buffer {
	var b Buffer
	return &b
}

// CreateBufferFromConfigFile ..
func CreateBufferFromConfigFile(path string) *Buffer {
	var b Buffer
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File error:", err)
	}
	err = yaml.Unmarshal(file, &b)
	if err != nil {
		fmt.Println("Unmarshall error:", err)
	}
	return &b
}

func (b *Buffer) SetBuffer(configs []Config) {
	for _, config := range configs {
		b.Configs = append(b.Configs, config)
	}
}

func (b *Buffer) GetService(serviceNameOrID interface{}) Config {
	switch serviceNameOrID.(type) {
	case string:
		for _, config := range b.Configs {
			if config.ServiceName == serviceNameOrID {
				return config
			}
		}
	case int:
		for _, config := range b.Configs {
			if config.Id == serviceNameOrID {
				return config
			}
		}
	default:
		fmt.Println("Type error")
	}
	return Config{}
}

func (b *Buffer) SetService(serviceNameOrID interface{}, newConfig Config) {
	switch serviceNameOrID.(type) {
	case string:
		for i, config := range b.Configs {
			if config.ServiceName == serviceNameOrID {
				b.Configs[i] = newConfig
			}
		}
	case int:
		for i, config := range b.Configs {
			if config.Id == serviceNameOrID {
				b.Configs[i] = newConfig
			}
		}
	default:
		fmt.Println("Type error")
	}
}

func (b *Buffer) DebugPrint() {
	for _, config := range b.Configs {
		fmt.Println(config)
	}
}
