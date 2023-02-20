package main

import (
	"flag"
	"main/grpc_server"
	. "main/logger"
)

var (
	LoggerConfigPath   string
	DatabaseConfigPath string
	ServerConfigPath   string
)

func init() {
	flag.StringVar(&LoggerConfigPath, "PathConfigLogger", "configs/logger-config.yaml", "Path to logger config file")
	flag.StringVar(&DatabaseConfigPath, "PathConfigDatabase", "configs/database-config.yaml", "Path to database config file")
	flag.StringVar(&ServerConfigPath, "PathConfigServer", "configs/server-config.yaml", "Path to server config file")
}

func main() {
	flag.Parse()
	logger := New()
	logger.Configurate(LoggerConfigPath)
	logger.Log.Info("Starting ...")

	var s = &grpc_server.GRPCServer{}
	s.Configurate(ServerConfigPath)
	s.Start()
}
