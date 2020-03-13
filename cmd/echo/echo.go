package main

import (
	"context"
	"github.com/nokamoto/envoy-external-authz/protobuf"
)

type Server struct{}

func (*Server) Echo(ctx context.Context, req *protobuf.EchoRequest) (*protobuf.EchoResponse, error) {
	return &protobuf.EchoResponse{
		Value: req.GetValue(),
	}, nil
}
