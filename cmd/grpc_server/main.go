package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/g1ltz0r/chat-server/cmd/helpers"

	desc "github.com/g1ltz0r/chat-server/pkg/chat_v1"
)

const grpcPort = 55556

type server struct {
	desc.UnimplementedChatV1Server
}

// Create chat
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create chat -> %+v", req)

	return &desc.CreateResponse{
		Id: helpers.GetRandID(),
	}, nil
}

// Delete chat
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete chat -> %+v", req)

	return &empty.Empty{}, nil
}

// SendMessage in chat
func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message -> %+v", req)

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
