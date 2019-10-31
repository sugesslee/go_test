package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"time"

	pb "gostudy/grpc/bidirectionalstream/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// ca 认证
	var tlsServerName = "client.io"
	certificate, err := tls.LoadX509KeyPair("/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/client.crt", "/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{certificate},
		ServerName:         tlsServerName, // NOTE: this is required!
		RootCAs:            certPool,
	})

	// set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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
