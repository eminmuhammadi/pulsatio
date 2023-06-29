package grpc

import (
	lib "github.com/eminmuhammadi/pulsatio/lib"
	grpc "google.golang.org/grpc"
	grpcInsecure "google.golang.org/grpc/credentials/insecure"
)

func Client(address string) (lib.PingPongClient, error) {
	// Create a new gRPC client connection
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(grpcInsecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Create a new ping pong client instance
	client := lib.NewPingPongClient(conn)

	return client, nil
}

func SecureClient(tlsFiles CertManager, address string, verify bool) (lib.PingPongClient, error) {
	creds, err := Credentials(tlsFiles, verify)
	if err != nil {
		return nil, err
	}

	// Dial the gRPC server with TLS credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	// Create a new ping pong client instance
	client := lib.NewPingPongClient(conn)

	return client, nil
}
