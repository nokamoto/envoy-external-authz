package main

import (
	"fmt"
	"github.com/nokamoto/envoy-external-authz/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	GrpcServerPort = "GRPC_SERVER_PORT"
)

func main() {
	port, err := strconv.Atoi(os.Getenv(GrpcServerPort))
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen tcp port (%d) - %v", port, err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	protobuf.RegisterEchoServer(server, &Server{})

	log.Printf("ready to serve %d", port)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("serve %v - %v", lis, err)
	}
}
