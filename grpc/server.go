package grpc

import (
	"net"

	lib "github.com/eminmuhammadi/pulsatio/lib"
	logger "github.com/eminmuhammadi/pulsatio/logger"
	grpc "google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	srv := grpc.NewServer()
	return srv
}

func GracefulStop(srv *grpc.Server) {
	srv.GracefulStop()
}

func NewSecureServer(tlsFiles CertManager, verify bool) (*grpc.Server, error) {
	creds, err := Credentials(tlsFiles, verify)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer(grpc.Creds(creds))

	return srv, nil
}

func RegisterServer(listener net.Listener, srv *grpc.Server) error {
	lib.RegisterPingPongServer(srv, &server{})

	logger.Printf("gRPC server listening on %s://%s", listener.Addr().Network(), listener.Addr().String())

	return srv.Serve(listener)
}

func StartServer(address string) error {
	// Create a new gRPC server instance
	srv := NewServer()

	// Create a TCP listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Register the server
	return RegisterServer(listener, srv)
}

func StartSecureServer(tlsFiles CertManager, address string, verify bool) error {
	// Create a new gRPC server instance
	srv, err := NewSecureServer(tlsFiles, verify)
	if err != nil {
		return err
	}

	logger.Printf("gRPC server is using tls files: %s, %s, %s", tlsFiles.CertFile, tlsFiles.KeyFile, tlsFiles.CAFile)

	// Create a TCP listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if verify {
		logger.Printf("gRPC server is using tls verification")
	} else {
		logger.Printf("gRPC server is not using tls verification")
	}

	// Register the server
	return RegisterServer(listener, srv)
}
