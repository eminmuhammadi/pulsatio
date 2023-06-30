package main

import (
	"fmt"
	"testing"

	cmd "github.com/eminmuhammadi/pulsatio/cmd"
	grpc "github.com/eminmuhammadi/pulsatio/grpc"
)

const tlsVerify = false

var port = 32000

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

func TestSecure(t *testing.T) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.SecureServer(serverCerts, address, tlsVerify)

	_, err := cmd.SecureClient(clientCerts, address, tlsVerify)

	if err != nil {
		t.Errorf("Secure test failed: %s", err)
	}
}

func TestInsecure(t *testing.T) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.Server(address)
	_, err := cmd.Client(address)

	if err != nil {
		t.Errorf("Insecure test failed: %s", err)
	}
}

func TestSecureServerInsecureClient(t *testing.T) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.SecureServer(serverCerts, address, tlsVerify)
	_, err := cmd.Client(address)

	if err == nil {
		t.Errorf("SecureServerInsecureClient test failed: %s", err)
	}
}

func TestInsecureServerSecureClient(t *testing.T) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.Server(address)
	_, err := cmd.SecureClient(clientCerts, address, tlsVerify)

	if err == nil {
		t.Errorf("InsecureServerSecureClient test failed: %s", err)
	}
}

func BenchmarkInsecure(b *testing.B) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.Server(address)
	for i := 0; i < b.N; i++ {
		_, err := cmd.Client(address)

		if err != nil {
			b.Errorf("Insecure test failed: %s", err)
		}
	}
}

func BenchmarkSecure(b *testing.B) {
	port++
	address := fmt.Sprintf("localhost:%d", port)

	go cmd.SecureServer(serverCerts, address, tlsVerify)
	for i := 0; i < b.N; i++ {
		_, err := cmd.SecureClient(clientCerts, address, tlsVerify)

		if err != nil {
			b.Errorf("Secure test failed: %s", err)
		}
	}
}
