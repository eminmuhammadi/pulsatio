package grpc

import (
	"context"

	config "github.com/eminmuhammadi/pulsatio/config"
	lib "github.com/eminmuhammadi/pulsatio/lib"
	logger "github.com/eminmuhammadi/pulsatio/logger"
	terminal "github.com/eminmuhammadi/pulsatio/terminal"
)

type server struct {
	lib.UnimplementedPingPongServer
}

func (s *server) Ping(ctx context.Context, req *lib.PingMessage) (*lib.PongMessage, error) {
	var pongMessage *lib.PongMessage

	msg := req.GetMessage()
	logger.Printf("Received: %s", msg)

	resp, err := terminal.Exec(msg, config.ServerTimeout)

	// Create and return the pong message
	if err != nil {
		pongMessage = &lib.PongMessage{Message: err.Error()}
	} else {
		pongMessage = &lib.PongMessage{Message: resp}
	}

	return pongMessage, nil
}

func Ping(client lib.PingPongClient, msg string) (*lib.PongMessage, error) {
	// Create a new context
	ctx := context.Background()

	return client.Ping(ctx, &lib.PingMessage{Message: msg})
}
