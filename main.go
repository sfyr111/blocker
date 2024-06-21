package main

import (
	"context"
	"fmt"
	"github.com/sfyr111/blocker/node"
	"github.com/sfyr111/blocker/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	node := node.NewNode()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port:", ":50051")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer client.Close()

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blocker-0.1",
		Height:  1,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
