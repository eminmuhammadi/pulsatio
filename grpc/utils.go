package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

type CertManager struct {
	CertFile string
	KeyFile  string
	CAFile   string
}

func Credentials(tlsFiles CertManager, verify bool) (credentials.TransportCredentials, error) {
	clientCert, err := tls.LoadX509KeyPair(tlsFiles.CertFile, tlsFiles.KeyFile)
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile(tlsFiles.CAFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: !verify,
	}

	creds := credentials.NewTLS(tlsConfig)

	return creds, nil
}
