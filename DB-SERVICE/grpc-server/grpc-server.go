package grpc_server

import (
	"main/grpc-server/grpc-server"
)

type GRPCServer struct {
	Config *grpc_server.Config
}

//New ..
func New() *GRPCServer {
	return &GRPCServer{
		Config: nil,
	}
}

//Start ..
func (s *GRPCServer) Start() error {
	return nil
}

//Configurate ..
func (s *GRPCServer) Configurate(path string) error {
	var err error
	s.Config, err = grpc_server.ParseConfigFile(path)
	return err
}
