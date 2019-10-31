package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"

	pb "gostudy/grpc/bidirectionalstream/proto"
)

const (
	address = "localhost:50051"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := c.Bidi(ctx)

	if err != nil {
		log.Fatalf("%v.BidiHello(_) = _, %v", c, err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note: %v", err)
			}
			fmt.Printf("AI: %s\n", in.Message)
		}
	}()

	for {
		request := &pb.HelloRequest{}
		_, _ = fmt.Scanln(&request.Name)

		if request.Name == "quit" {
			break
		}

		if err := stream.Send(request); err != nil {
			log.Fatalf("Failed to send a req: %v", err)
		}
	}

	_ = stream.CloseSend()
	<-waitc
}
