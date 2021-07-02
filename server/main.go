package main

import (
	"context"
	"log"
	"net"

	pb "github.com/haginom/grpcThree/chatfunc"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.MessageReply, error) {
	log.Fatalf("Recieved message body from client")
	return &pb.MessageReply{Body: "Hello from the Server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	//create a new gRPC server
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
