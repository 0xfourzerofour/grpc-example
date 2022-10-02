package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/joshpauline/grpc-example/usermgmt"
)

const (
	defaultName = "world"
	defaultAge  = 25
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name of user")
	age  = flag.Int("age", defaultAge, "Age of user")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.NewUser{Name: *name, Age: int32(*age)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Hello: %s", r.GetName())
}