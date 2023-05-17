package main

import (
	"github.com/common-nighthawk/go-figure"
	server "main/statistic_http_server"
)

func main() {
	serv := server.New()
	logo := figure.NewFigure("STATISTIC_SERVICE", "", true)
	logo.Print()
	serv.Start()
}
