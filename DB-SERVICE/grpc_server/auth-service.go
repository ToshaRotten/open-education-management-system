package grpc_server

import (
	"context"
	"google.golang.org/grpc"
	api "main/grpc_server/api/v1"
	"main/grpc_server/grpc_server_configurator"
	"net"
)

type GRPCServer struct {
	Config *grpc_server_configurator.Config
}

func (s *GRPCServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.UserReply, error) {
	return &api.UserReply{Message: "12312"}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedAuthServer() {
}

func (s *GRPCServer) Configurate(path string) error {
	var err error
	s.Config, err = grpc_server_configurator.ParseConfigFile(path)
	if err != nil {
		return err
	}
	return nil
}

func (s *GRPCServer) Start() error {
	grpc_s := grpc.NewServer()
	api.RegisterAuthServer(grpc_s, s)
	l, err := net.Listen("tcp", s.Config.Host+s.Config.Port)
	if err != nil {
		return err
	}
	err = grpc_s.Serve(l)
	if err != nil {
		return err
	}
	return nil
}
