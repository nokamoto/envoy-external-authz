package main

import (
	"context"
	md "github.com/nokamoto/envoy-external-authz/pkg/metadata"
	"github.com/nokamoto/envoy-external-authz/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	GrpcServerAddress = "GRPC_SERVER_ADDRESS"
	APIKey            = "APIKEY"
)

func main() {
	addr := os.Getenv(GrpcServerAddress)
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("ready to %v", addr)

	echo := protobuf.NewEchoClient(cc)

	i := 0
	for {
		req := &protobuf.EchoRequest{
			Value: strconv.Itoa(i),
		}

		log.Printf("req=%v", req)

		ctx := metadata.AppendToOutgoingContext(context.Background(), md.XMyAPIKey, os.Getenv(APIKey))

		res, err := echo.Echo(ctx, req)

		log.Printf("res=%v, err=%v", res, err)

		time.Sleep(3 * time.Second)

		i += 1
	}
}
