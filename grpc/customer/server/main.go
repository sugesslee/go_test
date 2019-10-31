package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gostudy/grpc/customer/proto"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.CustomerRequest
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	// 证书认证
	//creds, err := credentials.NewServerTLSFromFile("/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/server.crt", "/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/server.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	// Creates a new gRPC server

	// ca认证
	certificate, err := tls.LoadX509KeyPair("/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/server.crt", "/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("/Users/red/Desktop/workspace/project/go-project/gostudy/grpc/key/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCustomerServer(s, &server{})
	_ = s.Serve(lis)
}
