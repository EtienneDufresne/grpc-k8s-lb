package main

import (
	"log"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/EtienneDufresne/grpc-k8s-lb/protos"
)

type server struct {
	Host string
}

func NewServer(host string) *server {
	return &server{
		Host: host,
	}
}

func (s *server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, s)

	go grpcServer.Serve(listener)

	log.Printf("Server listening on %s", host)

	<-ctx.Done()

	return nil
}

func (s *server) EchoMessage(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("=================================================================")
	log.Printf("Received %s", in.Message)
	log.Printf("Sending %s", in.Message)
	serverID := os.Getenv("HOSTNAME")
	if serverID == "" {
		serverID = s.Host
	}
	return &pb.EchoResponse{
		Message:  in.Message,
		ServerID: serverID,
	}, nil
}
