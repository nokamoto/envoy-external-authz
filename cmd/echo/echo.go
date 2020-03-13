package main

import (
	"context"
	"github.com/nokamoto/envoy-external-authz/protobuf"
	"google.golang.org/grpc/metadata"
	"log"
)

type Server struct{}

func (*Server) Echo(ctx context.Context, req *protobuf.EchoRequest) (*protobuf.EchoResponse, error) {
	md, found := metadata.FromIncomingContext(ctx)
	log.Printf("metadata=%v, found=%v", md, found)

	return &protobuf.EchoResponse{
		Value: req.GetValue(),
	}, nil
}
