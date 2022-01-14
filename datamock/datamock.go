package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/MayankPandey01/GetUserByName"
	"google.golang.org/grpc"
)

const (
	port = ":10000"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) GetMockUserData(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var name string = in.GetName()
	if len(name) < 6 {
		//log.Panic("Name " + name + " is Less than 6 Characters")
		return nil, errors.New("Name " + name + " is Less than 6 Characters")
	}
	var class int = len(name)
	var roll int = len(name) * 10
	return &pb.User{Name: name, Class: int64(class), Roll: int32(roll)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Mock Data Service Started")
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
