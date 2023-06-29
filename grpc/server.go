package grpc

import (
	"log"
	"net"

	lib "github.com/eminmuhammadi/pulsatio/lib"
	grpc "google.golang.org/grpc"
)

type server struct {
	lib.UnimplementedPingPongServer
}

func NewServer() *grpc.Server {
	srv := grpc.NewServer()
	return srv
}

func RegisterServer(listener net.Listener, srv *grpc.Server) error {
	lib.RegisterPingPongServer(srv, &server{})

	log.Printf("gRPC server listening on %s://%s", listener.Addr().Network(), listener.Addr().String())

	return srv.Serve(listener)
}

func StartServer(protocol string, address string) error {
	// Create a new gRPC server instance
	srv := NewServer()

	// Create a TCP listener
	listener, err := net.Listen(protocol, address)
	if err != nil {
		return err
	}

	// Register the server
	return RegisterServer(listener, srv)
}
