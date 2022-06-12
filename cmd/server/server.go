package main

import (
	"log"
	"net"

	"github.com/LucasGois1/learning-grpc/pb"
	"github.com/LucasGois1/learning-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main () {

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen server", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to create server", err)
	}
}
