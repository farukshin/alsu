package main

import (
	"context"
	"log"
	"net"

	pb "farukshin.com/alsu/proto"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const (
	port           = ":50052"
	orderBatchSize = 3
)

type server struct {
	str1 string
}

func (s *server) PingServer(ctx context.Context, req *wrapper.StringValue) (*wrapper.StringValue, error) {
	return &wrapper.StringValue{Value: "pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAlsuServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
