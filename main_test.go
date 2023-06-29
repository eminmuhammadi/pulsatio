package main

import (
	"testing"

	grpc "github.com/eminmuhammadi/pulsatio/grpc"
	lib "github.com/eminmuhammadi/pulsatio/lib"
)

var serverCerts = grpc.CertManager{
	CertFile: ".tls/server-cert.pem",
	KeyFile:  ".tls/server-key.pem",
	CAFile:   ".tls/ca-cert.pem",
}

var clientCerts = grpc.CertManager{
	CertFile: ".tls/client-cert.pem",
	KeyFile:  ".tls/client-key.pem",
	CAFile:   ".tls/ca-cert.pem",
}

func SecureServer(tlsFiles grpc.CertManager, address string, verify bool) {
	err := grpc.StartSecureServer(tlsFiles, address, verify)
	if err != nil {
		panic(err)
	}
}

func SecureClient(tlsFiles grpc.CertManager, address string, verify bool) (*lib.PongMessage, error) {
	client, err := grpc.SecureClient(tlsFiles, address, verify)
	if err != nil {
		panic(err)
	}

	return grpc.Ping(client)
}

func Server(address string) {
	err := grpc.StartServer(address)
	if err != nil {
		panic(err)
	}
}

func Client(address string) (*lib.PongMessage, error) {
	client, err := grpc.Client(address)
	if err != nil {
		panic(err)
	}

	return grpc.Ping(client)
}

func TestSecure(t *testing.T) {
	address := "localhost:32001"
	verify := false

	go SecureServer(serverCerts, address, verify)

	_, err := SecureClient(clientCerts, address, verify)

	if err != nil {
		t.Errorf("Secure test failed: %s", err)
	}
}

func TestInsecure(t *testing.T) {
	address := "localhost:32002"

	go Server(address)
	_, err := Client(address)

	if err != nil {
		t.Errorf("Insecure test failed: %s", err)
	}
}

func TestSecureServerInsecureClient(t *testing.T) {
	address := "localhost:32003"
	verify := false

	go SecureServer(serverCerts, address, verify)
	_, err := Client(address)

	if err == nil {
		t.Errorf("SecureServerInsecureClient test failed: %s", err)
	}
}

func TestInsecureServerSecureClient(t *testing.T) {
	address := "localhost:32004"

	go Server(address)
	_, err := SecureClient(clientCerts, address, true)

	if err == nil {
		t.Errorf("InsecureServerSecureClient test failed: %s", err)
	}
}
