package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gostudy/grpc/serversidestream/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement greeterServer
type server struct {
}

func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	for idx := 0; idx < 10; idx++ {
		log.Printf("recv message:Name: %s, idx: %d", in.Name, idx)
		_ = stream.Send(&pb.HelloReply{Message: fmt.Sprintf("Hello %s %d", in.Name, idx)})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC  server.
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
