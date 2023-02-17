package main

import (
	"flag"
	. "main/loger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configurator-path", "configs/apiserver.yaml", "path to configurator file")
}

func main() {
	flag.Parse()
	GlobalLoger.Info("Starting ...")
}
