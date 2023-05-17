package main

import (
	"github.com/common-nighthawk/go-figure"
	server "main/schedule_http_server"
)

func main() {
	s := server.New()
	logo := figure.NewFigure("SCHEDULE MODULE", "", true)
	logo.Print()
	s.Start()
}
