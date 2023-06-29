#!/bin/env sh

# Check if protoc is installed
if ! [ -x "$(command -v protoc)" ]; then
    echo 'Error: protoc is not installed.' >&2
    # Install protoc using golang
    go get -u github.com/golang/protobuf/protoc-gen-go@latest
    go install github.com/golang/protobuf/protoc-gen-go@latest
fi

# Check if protoc-gen-go-grpc is installed
if ! [ -x "$(command -v protoc-gen-go-grpc)" ]; then
    echo 'Error: protoc-gen-go-grpc is not installed.' >&2
    # Install protoc-gen-go-grpc using golang
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Generate gRPC code
protoc -I=proto --go_out=. --go-grpc_out=. proto/*.proto
