package logger

import (
	"github.com/sirupsen/logrus"
	logger_configurator "main/logger/logger-configurator"
)

type globalLoger struct {
	Config *logger_configurator.Config
	Log    *logrus.Logger
}

//New ..
func New() *globalLoger {
	var g globalLoger
	g.Log = logrus.New()
	return &g
}

// Configure
func (g *globalLoger) Configurate(path string) (*globalLoger, error) {
	var err error
	g.Config, err = logger_configurator.ParseConfigFile(path)
	if err != nil {
		return nil, err
	}
	return g, err
}
