package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MayankPandey01/GetUserByName"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
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

	var new_users = []string {"Mayank"} // Enter Your Names Here [INPUT SECTION]

	for _, name := range new_users {
		r, err := c.GetUserByName(ctx, &pb.NewUser{Name: name})
		//r, err := c.GetMockUserData(ctx, &pb.NewUser{Name: name})
		if err != nil {
			log.Fatalf("could not Get user: %v", err)
		}
		log.Printf(`
:::User Details:::
NAME: %s
CLASS: %d
ROLL: %d`, r.GetName(), r.GetClass(), r.GetRoll())

	}
}