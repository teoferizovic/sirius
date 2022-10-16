package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/teo/sirius/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 30

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
		NAME: %s
		AGE: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())

	}

	params := &pb.GetUserParam{}

	r, err := c.GetUsers(ctx, params)

	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}

	log.Print("\nUSER LIST: \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())

	//-------------------------------
	r2, err := c.GetLastUser(ctx, params)

	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}

	//log.Print("\LAST nUSER: \n")
	fmt.Println(r2.GetName())

	//-------------------------------
	log.Print("\nSAY HELLO: \n")

	r3, err := c.SayHello(ctx, &pb.HelloMessage{Msg: "teo is here"})

	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}

	fmt.Println(r3)
	//-------------------------------
	log.Print("\nSAY HELLO2: \n")

	r4, err := c.SayHello2(ctx, &pb.HelloMessage2{Msg: "teo is here from hello", Number: 34})

	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}

	fmt.Println(r4)
	//-------------------------------
}
