package main

import (
	"log"
	"net"

	pb "Week04/api/article/hw04/v1"
	"Week04/internal/hw04/v1/biz"
	"Week04/internal/hw04/v1/service"
	"Week04/internal/hw04/v1/data"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	data := data.NewHWRepo()
	uc := biz.NewHWUsecase(data)
	svc := service.NewHWService(uc)

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, svc)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
