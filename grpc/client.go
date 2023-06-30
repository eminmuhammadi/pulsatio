package grpc

import (
	lib "github.com/eminmuhammadi/pulsatio/lib"
	grpc "google.golang.org/grpc"
	grpcInsecure "google.golang.org/grpc/credentials/insecure"
)

func Connect(address string) (*grpc.ClientConn, error) {
	// Create a new gRPC client connection
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(grpcInsecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Client(conn *grpc.ClientConn) lib.PingPongClient {
	return lib.NewPingPongClient(conn)
}

func SecureConnect(tlsFiles CertManager, address string, verify bool) (*grpc.ClientConn, error) {
	creds, err := Credentials(tlsFiles, verify)
	if err != nil {
		return nil, err
	}

	// Dial the gRPC server with TLS credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
