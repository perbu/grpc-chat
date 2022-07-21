package chat

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Send(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Got message: %s", message.Content)
	ret := &Message{
		Content: fmt.Sprintf("Message: %s", message.Content),
	}
	return ret, nil
}
