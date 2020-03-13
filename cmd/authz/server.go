package main

import (
	"context"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/nokamoto/envoy-external-authz/pkg/metadata"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"log"
)

type Server struct {
	apikey string
}

func (s *Server) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	http := req.GetAttributes().GetRequest().GetHttp()

	log.Printf("path=%v", http.GetPath())
	log.Printf("metadata=%v", http.GetHeaders())

	unauthorized := auth.CheckResponse{
		Status: &status.Status{
			Code:    int32(code.Code_PERMISSION_DENIED),
			Message: "external authorization failed",
		},
		HttpResponse: nil,
	}

	apikey, found := http.GetHeaders()[metadata.XMyAPIKey]
	if !found {
		log.Printf("no metadata")
		return &unauthorized, nil
	}

	if s.apikey != apikey {
		log.Printf("%v != %v", s.apikey, apikey)
		return &unauthorized, nil
	}

	return &auth.CheckResponse{}, nil
}
