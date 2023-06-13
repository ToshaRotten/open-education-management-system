package configurator

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator/models"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Configurator struct {
	Buffer *buffer.Buffer
	Path   string
}

func New(path string) *Configurator {
	var c Configurator
	c.Buffer = buffer.New()
	c.Path = path
	c.Buffer.SetBuffer(c.parseFile())
	return &c
}

// ParseFile ..
func (c *Configurator) parseFile() []models.Config {
	var configs models.Configs
	file, err := ioutil.ReadFile(c.Path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = yaml.Unmarshal(file, &configs)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return configs.Configs
}
