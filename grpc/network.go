package grpc

import "net"

func Listen(protocol string, address string) (net.Listener, error) {
	return net.Listen(protocol, address)
}
