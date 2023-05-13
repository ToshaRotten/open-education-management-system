package configurator

import (
	"main/configurator/buffer"
)

type Configurator struct {
	Buffer *buffer.Buffer
}

func New() *Configurator {
	var c Configurator
	c.Buffer = buffer.New()
	return &c
}
