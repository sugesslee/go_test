package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	pb "gostudy/grpc/bidirectionalstream/proto"
	"io"
	"io/ioutil"
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

	var tlsServerName = "server.io"
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

	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC  server.
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
