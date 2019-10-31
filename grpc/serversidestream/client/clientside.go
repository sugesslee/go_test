package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"

	pb "gostudy/grpc/serversidestream/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// contact the server and print out its response.
	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.LotsOfReplies() = _, %v", c, err)
		}

		log.Printf("Greeting: %s\n", reply.Message)
	}
}
