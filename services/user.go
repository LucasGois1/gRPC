package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/LucasGois1/learning-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Printf("Saving on database...")

	return &pb.User{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResponseStream{
		Status: "Initializing",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResponseStream{
		Status: "Saving on database",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResponseStream{
		Status: "Completed",
		User: &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	users := []*pb.User{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				Users: users,
			})
		}

		if err != nil {
			log.Fatalf("Error: %v", err)
			return err
		}

		users = append(users, req)

	}
}
