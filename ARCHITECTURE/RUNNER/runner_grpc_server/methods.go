package runner_grpc_server

import (
	"context"
	pb2 "github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/RUNNER/runner_grpc_server/proto/pb"
)

type RunnerServer struct {
	pb2.RunnerServer
}

func (s *RunnerServer) RunService(ctx context.Context, in *pb2.RunRequest) (*pb2.RunResponse, error) {
	//ServiceNamesToStart := in.GetServiceNames()
	//var ServicesToStart []service.Service
	return nil, nil
}
