package grpc

import (
	"context"

	lib "github.com/eminmuhammadi/pulsatio/lib"
	logger "github.com/eminmuhammadi/pulsatio/logger"
)

func (s *server) Ping(ctx context.Context, req *lib.PingMessage) (*lib.PongMessage, error) {
	message := req.GetMessage()
	logger.Printf("Received ping message: %s", message)

	// Create and return the pong message
	pongMessage := &lib.PongMessage{Message: "Pong!"}
	return pongMessage, nil
}

func Ping(client lib.PingPongClient) (*lib.PongMessage, error) {
	// Create a new context
	ctx := context.Background()

	return client.Ping(ctx, &lib.PingMessage{Message: "Ping!"})
}
