package main

import (
	"context"
	"log"
	"time"

	pb "farukshin.com/alsu-client/proto"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAlsuServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	retrievedOrder, err := client.PingServer(ctx, &wrapper.StringValue{Value: "ping"})
	log.Print("PingServer() Response -> : ", retrievedOrder)
}
