package main

import (
	"flag"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/configurator_http_client"
	"github.com/common-nighthawk/go-figure"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	configPath   string
	logger       = logrus.New()
	configurator = configurator_http_client.New()
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/config.yaml", "Path to server config file")
}

func main() {
	flag.Parse()
	logger.Info("Parse flags ...")
	logo := figure.NewFigure("AUTH SERVICE", "", true)
	logo.Print()
	logger.Info("Starting ...")

	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("Trying to get config from CONFIGURATOR_SERVICE"),
	)

	for i := 0; i < 100; i++ {
		bar.Add(20)
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()

	//Check file

	//Check configurator service
	config, err := configurator.GetConfig()
	if err != nil {
		logger.Warn(err)
		logger.Info("Set a default config")
	}

	config.Host = "localhost"
	config.Port = ":8080"

	server := auth_http_server.New()
	if err := server.Start(&config); err != nil {
		logger.Error(err)
	}
}
