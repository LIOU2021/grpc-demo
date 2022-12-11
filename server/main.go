package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pbUsers "github.com/LIOU2021/grpc-demo/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UsersServer struct {
	pbUsers.UnimplementedUsersServer
}

func (server *UsersServer) Greet(ctx context.Context, req *pbUsers.GreetRequest) (rep *pbUsers.GreetResponse, err error) {
	rep = &pbUsers.GreetResponse{}
	data := []*pbUsers.User{
		{
			Age:    10,
			Name:   "f1",
			Gender: pbUsers.Gender_female,
		},
		{
			Age:    11,
			Name:   "f2",
			Gender: pbUsers.Gender_female,
		},
	}

	data = append(data, &pbUsers.User{
		Age:    12,
		Name:   "m1",
		Gender: pbUsers.Gender_male,
	})

	rep.Data = data
	rep.Page = req.Page
	rep.Limit = req.Limit
	return
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbUsers.RegisterUsersServer(s, &UsersServer{})
	//you can register multiple grpc server here
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
