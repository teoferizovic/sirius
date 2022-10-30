package main

import (
	"context"
	"log"

	//"math/rand"
	"net"

	userpb "github.com/teo/sirius/proto"
	"google.golang.org/grpc"
)

const (
	port = "5055"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     req.Uuid,
			FullName: "Teo",
		},
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServer := &userService{}

	userpb.RegisterUserServiceServer(grpcServer, userServer)

	grpcServer.Serve(lis)
}
