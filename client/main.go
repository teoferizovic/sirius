package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/teo/sirius/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:5055"
)

func main() {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "1-2-3",
	})

	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("%+v\n", res)

	res2, err := client.GetClient(context.Background(), &userpb.GetClientRequest{
		FirstName: "John",
		LastName:  "Doe",
	})

	if err != nil {
		log.Fatalf("fail to GetClient: %v", err)
	}

	fmt.Printf("%+v\n", res2)
}
