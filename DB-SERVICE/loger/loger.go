package loger

import (
	"github.com/sirupsen/logrus"
	"main/loger/configurator"
)

type globalLoger struct {
	Config *configurator.Config
	Loger  *logrus.Logger
}

//init
func Init() *globalLoger {
	var g globalLoger
	g.Configurate()
	return &g
}

func (g *globalLoger) Configurate() *globalLoger {
	return g
}
