package main

import (
	"context"
	"fmt"
	"io"

	"github.com/LucasGois1/learning-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUserVerbose(client)
}

func GetUser(client pb.UserServiceClient) {
	req := pb.User{
		Id:    1,
		Name:  "Lucas",
		Email: "lucas@mail.com",
	}

	response, err := client.AddUserVerbose(context.Background(), &req)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", response)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := pb.User{
		Id:    1,
		Name:  "Lucas",
		Email: "lucas2@mail.com.br",
	}

	stream, err := client.AddUserVerbose(context.Background(), &req)

	if err != nil {
		panic(err)
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("Response: %v\n", response)
	}
}
