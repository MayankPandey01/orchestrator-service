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
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) GetUserByName(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	return nil, errors.New("not implemented yet. Mayank will implement me")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
