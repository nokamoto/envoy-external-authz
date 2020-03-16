package main

import (
	"context"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
		// https://github.com/envoyproxy/envoy/blob/e1962d76c073bb48c57acd0ff2b57390a22394b7/api/envoy/service/auth/v2/external_auth.proto#L38-L50
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

	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		// https://github.com/envoyproxy/envoy/blob/e1962d76c073bb48c57acd0ff2b57390a22394b7/api/envoy/service/auth/v2/external_auth.proto#L53-L61
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Header: &core.HeaderValue{
							Key:   metadata.XMyUsername,
							Value: "foo",
						},
						Append: &wrappers.BoolValue{
							Value: false,
						},
					},
				},
			},
		},
	}, nil
}
