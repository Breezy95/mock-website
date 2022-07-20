package proto

import(
	"log"
	"context"
	
)

type Server struct {
	// Embed the unimplemented server
UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Server says Hi"}, nil

}