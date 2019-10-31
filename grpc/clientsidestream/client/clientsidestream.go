package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	pb "gostudy/grpc/clientsidestream/proto"
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

	stream, err := c.LotsOfReplies(ctx)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for idx := 0; idx < 10; idx++ {
		if err := stream.Send(&pb.HelloRequest{
			Name:                 name,
			Index:                int32(idx),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}); err != nil {
			log.Fatalf("send error: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got  error %v, want %v", stream, err, nil)
	}

	log.Printf("Greeting: %s\n", reply.Message)
}
