package runner_http_server

import (
	"context"
	"main/ServiceController/service"
	pb2 "main/runner_server/proto/pb"
)

type RunnerServer struct {
	pb2.RunnerServer
}

func (s *RunnerServer) RunService(ctx context.Context, in *pb2.RunRequest) (*pb2.RunResponse, error) {
	ServiceNamesToStart := in.GetServiceNames()
	var ServicesToStart []service.Service
	return nil, nil
}
