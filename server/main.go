package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/teo/sirius/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		users_list: &pb.UserList{},
		last_user:  &pb.User{},
		//hello_msg:  &pb.HelloMessage{},
		//hello_msg2: &pb.HelloMessage2{},
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	users_list *pb.UserList
	last_user  *pb.User
	//hello_msg  *pb.HelloMessage
	//hello_msg2 *pb.HelloMessage2
}

func (server *UserManagementServer) Run() error {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)

	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id = int32(rand.Intn(100))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	server.users_list.Users = append(server.users_list.Users, created_user)
	server.last_user = created_user
	return created_user, nil
}

func (server *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUserParam) (*pb.UserList, error) {
	return server.users_list, nil
}

func (server *UserManagementServer) GetLastUser(ctx context.Context, in *pb.GetUserParam) (*pb.User, error) {
	return server.last_user, nil
}

func (server *UserManagementServer) SayHello(ctx context.Context, msg *pb.HelloMessage) (*pb.HelloMessage, error) {
	var sayHello *pb.HelloMessage
	sayHello = &pb.HelloMessage{Msg: msg.GetMsg()}
	return sayHello, nil
}

func (server *UserManagementServer) SayHello2(ctx context.Context, msg *pb.HelloMessage2) (*pb.HelloMessage2, error) {
	var kk *pb.HelloMessage2
	kk = &pb.HelloMessage2{Msg: msg.GetMsg(), Number: msg.GetNumber()}
	return kk, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
