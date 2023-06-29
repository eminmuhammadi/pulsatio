package main

import (
	grpc "github.com/eminmuhammadi/pulsatio/grpc"
)

const protocol = "tcp"
const address = "localhost:8080"

func Server() {
	err := grpc.StartServer(protocol, address)
	if err != nil {
		panic(err)
	}
}

func Client() {
	client, err := grpc.Client(address)
	if err != nil {
		panic(err)
	}

	grpc.Ping(client)
}

func main() {
	go Server()
	Client()
}
