package grpc

import (
	"context"
	"log"

	lib "github.com/eminmuhammadi/pulsatio/lib"
)

func (s *server) Ping(ctx context.Context, req *lib.PingMessage) (*lib.PongMessage, error) {
	message := req.GetMessage()
	log.Printf("Received ping message: %s", message)

	// Create and return the pong message
	pongMessage := &lib.PongMessage{Message: "Pong!"}
	return pongMessage, nil
}

func Ping(client lib.PingPongClient) {
	// Create a new context
	ctx := context.Background()

	// Ping the server and print the response
	pongResponse, err := client.Ping(ctx, &lib.PingMessage{Message: "Ping!"})
	if err != nil {
		log.Fatalf("Error when calling Ping: %v", err)
	}

	log.Printf("Response from server: %s", pongResponse.GetMessage())
}
