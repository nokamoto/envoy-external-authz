package main

import (
	"context"
	"github.com/nokamoto/envoy-external-authz/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	GrpcServerAddress = "GRPC_SERVER_ADDRESS"
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

		res, err := echo.Echo(context.Background(), req)

		log.Printf("res=%v, err=%v", res, err)

		time.Sleep(3 * time.Second)

		i += 1
	}
}
