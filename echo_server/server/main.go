package main

import (
	"context"

	"learn.trpc.go/echo_server/pb"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

type EchoServer struct{}

func main() {
	server := trpc.NewServer()
	pb.RegisterEchoServerService(server, &EchoServer{})
	server.Serve()
}

func (s *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Infof("Received: %v", req.Msg)
	return &pb.EchoReply{Msg: req.Msg}, nil
}