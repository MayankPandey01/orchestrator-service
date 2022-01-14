package main

import (
	"context"
	"log"
	"net"
	"time"
	"errors"

	pb "github.com/MayankPandey01/GetUserByName"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}




func (s *UserManagementServer) GetUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	var address = "localhost:10000"
	log.Printf("Received: %v", in.GetName())
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//sending Data to GetMockData Method
	var name_send=in.GetName()
	r, err := c.GetMockUserData(ctx1, &pb.NewUser{Name: name_send})
	
	if err != nil {
		//log.Fatalf("\ncould not Get user: %v", err)
		return nil, errors.New("Name " + name_send + " is Less than 6 Characters")
	}
	//Parsing and Returning Data Back To client
	var name = r.GetName()
	var class= r.GetClass()
	var roll= r.GetRoll()
	return &pb.User{Name: name, Class: int64(class), Roll: int32(roll)}, err
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Orchestrator 2 Started")
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}	
}
