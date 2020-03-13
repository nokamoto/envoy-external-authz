PROTOFILES = $(wildcard protobuf/*.proto)
PBGOFILES = $(PROTOFILES:.proto=.pb.go)

all: go.mod $(PBGOFILES)
	go fmt ./cmd/*
	docker-compose build

$(PBGOFILES): $(PROTOFILES)
	prototool format -w protobuf
	protoc -I . protobuf/*.proto --go_out=plugins=grpc,paths=source_relative:.

go.mod:
	go mod init github.com/nokamoto/envoy-external-authz
