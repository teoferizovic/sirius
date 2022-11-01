package main

import (
	"context"

	userpb "github.com/teo/sirius/proto"
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

func (u *userService) GetClient(_ context.Context, req *userpb.GetClientRequest) (*userpb.GetClientResponse, error) {
	return &userpb.GetClientResponse{
		FullName: req.GetFirstName() + " " + req.GetLastName(),
	}, nil
}
