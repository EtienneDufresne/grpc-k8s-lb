package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/EtienneDufresne/grpc-k8s-lb/protos"
)

type server struct {
	Host           string
	savedCustomers []*pb.CustomerRequest
}

func NewServer(host string) *server {
	return &server{
		Host: host,
	}
}

func (s *server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	gs := grpc.NewServer()
	pb.RegisterCustomerServer(gs, s)

	go gs.Serve(lis)

	<-ctx.Done()

	return nil
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	log.Printf("CreateCustomer %s", in.Name)
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	log.Printf("GetCustomers")
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
