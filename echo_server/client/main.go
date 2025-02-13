package main

import (
	"context"

	"learn.trpc.go/echo_server/pb"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	c := pb.NewEchoServerClientProxy(client.WithTarget("ip://127.0.0.1:8000"))
	rsp, err := c.Echo(context.Background(), &pb.EchoRequest{Msg: "hello"})
	if err != nil {
		log.Error(err)
	}
	log.Info(rsp.Msg)
}
