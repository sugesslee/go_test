package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gostudy/grpc/bidirectionalstream/proto"
	"io"
	"log"
	"net"
	"strings"
)

const (
	port = ":50051"
)

// server is used to implement greeterServer
type server struct {
}

func (s *server) Bidi(stream pb.Greeter_BidiServer) error {
	for {
		in, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		message := strings.Replace(in.Name, "吗", "", -1)
		message = strings.Replace(message, "?", "!", -1)

		err = stream.Send(&pb.HelloReply{Message: message})
		//log.Printf("Send: Hello %s, total %d", name, total)
		if err != nil {
			return err
		}
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
