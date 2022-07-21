package main

import (
	"context"
	"github.com/perbu/grpc-chat/chat"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Content: "Hello from the client!",
	}
	start := time.Now()
	response, err := c.Send(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Time taken: %v", time.Since(start))
	log.Printf("Response from Server: %s", response.Content)

}
