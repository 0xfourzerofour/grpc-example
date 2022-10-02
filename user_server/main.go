package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/joshpauline/grpc-example/usermgmt"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserManagementServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {

	id := int32(12345678)

	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: id}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
