package main

import (
	"context"
	"log"
	"time"

	pb "github.com/haginom/grpcThree/chatfunc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c.SayHello(ctx, &pb.Message{Body: "Hello from the Client"})

	r, err := c.SayHello(ctx, &pb.Message{Body: "Hello from the Client"})
	if err != nil {
		log.Printf("Error when calling SayHello: %v", err)
	}
	log.Printf("Response from server: %s", r.Body)
}
