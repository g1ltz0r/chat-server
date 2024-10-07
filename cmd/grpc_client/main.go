package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/g1ltz0r/chat-server/cmd/helpers"

	desc "github.com/g1ltz0r/chat-server/pkg/chat_v1"
)

const address = "127.0.0.1:55556"

func createChat(ctx context.Context, c desc.ChatV1Client) {
	r, err := c.Create(ctx, &desc.CreateRequest{
		Usernames: []string{gofakeit.Name(), gofakeit.Name(), gofakeit.Name()},
	})
	if err != nil {
		log.Fatalf("failed to create chat: %v", err)
	}

	log.Printf(color.RedString("Create Chat:\n"), color.GreenString("%+v", r))
}

func deleteChat(ctx context.Context, c desc.ChatV1Client) {
	r, err := c.Delete(ctx, &desc.DeleteRequest{Id: helpers.GetRandID()})
	if err != nil {
		log.Fatalf("failed to delete chat: %v", err)
	}

	log.Printf(color.RedString("Delete Chat:\n"), color.GreenString("%+v", r))
}

func sendMsg(ctx context.Context, c desc.ChatV1Client) {
	r, err := c.SendMessage(ctx, &desc.SendMessageRequest{
		From:      gofakeit.Name(),
		Text:      gofakeit.HackerPhrase(),
		Timestamp: timestamppb.New(gofakeit.Date()),
	})
	if err != nil {
		log.Fatalf("failed to delete chat: %v", err)
	}

	log.Printf(color.RedString("Delete Chat:\n"), color.GreenString("%+v", r))
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close the connection: %v", err.Error())
		}
	}()

	c := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createChat(ctx, c)
	deleteChat(ctx, c)
	sendMsg(ctx, c)
}
