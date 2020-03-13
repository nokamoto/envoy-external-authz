package main

import (
	"fmt"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	GrpcServerPort = "GRPC_SERVER_PORT"
	APIKey         = "APIKEY"
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

	auth.RegisterAuthorizationServer(server, &Server{
		apikey: os.Getenv(APIKey),
	})

	log.Printf("ready to serve %d", port)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("serve %v - %v", lis, err)
	}
}
