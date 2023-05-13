package main

import (
	"fmt"
	"google.golang.org/grpc"
	"main/runner_server/proto/pb"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()

	pb.RegisterRunnerServer(s, &runner_http_server.RunnerServer{})
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
