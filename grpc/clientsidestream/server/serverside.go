package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gostudy/grpc/clientsidestream/proto"
	"io"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement greeterServer
type server struct {
}

func (s *server) LotsOfReplies(stream pb.Greeter_LotsOfRepliesServer) error {
	var total int32
	var name string

	for {
		greeting, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{
				Message:              fmt.Sprintf("Hello %s, total %d", name, total),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})
		}
		log.Printf("Send: Hello %s, total %d", name, total)
		if err != nil {
			return err
		}
		name = greeting.Name
		total += greeting.Index
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
