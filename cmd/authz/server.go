package main

import (
	"context"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"log"
)

type Server struct{}

func (*Server) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	http := req.GetAttributes().GetRequest().GetHttp()

	log.Printf("path=%v", http.GetPath())
	log.Printf("metadata=%v", http.GetHeaders())

	return &auth.CheckResponse{}, nil
}
