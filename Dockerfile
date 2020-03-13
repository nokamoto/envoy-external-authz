FROM golang:1.13.8-alpine3.11 AS build

RUN apk update && apk add git

WORKDIR /src/cmd

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

ARG cmd

COPY protobuf protobuf
COPY cmd/$cmd cmd/$cmd

RUN go install ./cmd/$cmd

FROM alpine:3.11

RUN apk update && apk add --no-cache ca-certificates

ARG cmd

COPY --from=build /go/bin/$cmd /usr/local/bin/cmd

ENTRYPOINT [ "cmd" ]
