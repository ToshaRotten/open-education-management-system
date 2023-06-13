package main

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator_http_server"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	logo := figure.NewFigure("CONFIGURATOR_SERVICE", "", true)
	logo.Print()
	serv := configurator_http_server.New()
	serv.Start()
}
