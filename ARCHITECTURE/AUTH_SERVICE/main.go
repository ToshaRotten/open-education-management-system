package main

import (
	"flag"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/configurator_http_client"
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	configPath string
	host       string
	port       string
	logger     = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			ForceColors: true,
			ForceQuote:  true,
		},
	}
	configurator = configurator_http_client.New()
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/config.yaml", "Path to server config file")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.StringVar(&port, "port", ":9906", "port")
}

func main() {
	flag.Parse()
	logo := figure.NewFigure("AUTH SERVICE", "", true)
	logo.Print()

	logger.Info(fmt.Sprintf("Parse flags, flag values: host: %s, port: %s", host, port))
	logger.Info("Starting ...")
	//bar := progressbar.NewOptions(1000,
	//	progressbar.OptionSetWidth(15),
	//	progressbar.OptionSetDescription("Trying to get config from CONFIGURATOR_SERVICE"),
	//)

	//for i := 0; i < 100; i++ {
	//	bar.Add(20)
	//	time.Sleep(40 * time.Millisecond)
	//}
	//fmt.Println()

	//Check file

	//Check configurator service
	config, err := configurator.GetConfig("AUTH_SERVICE")
	if err != nil {
		logger.Warn(err)
		logger.Info("Set a default config")
		config.Host = host
		config.Port = port
	}

	server := auth_http_server.New()
	if err := server.Start(&config); err != nil {
		logger.Error(err)
	}
}
