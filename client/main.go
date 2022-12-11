package main

import (
	"context"
	"fmt"
	"log"

	pbUsers "github.com/LIOU2021/grpc-demo/proto/users"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pbUsers.NewUsersClient(conn)
	feature, err := client.Greet(
		context.Background(),
		&pbUsers.GreetRequest{
			Limit: 7,
			Page:  21,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(feature)
}
