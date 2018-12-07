package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/EtienneDufresne/grpc-k8s-lb/protos"
)

type client struct {
	Host string
}

func NewClient(host string) *client {
	return &client{
		Host: host,
	}
}

func (c *client) Run(ctx context.Context) error {
	conn, err := grpc.DialContext(ctx, c.Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", c.Host, err)
		return err
	}
	defer conn.Close()

	log.Printf("Client connecting to %s", host)

	go sendEchoMessagePeriodically(ctx, conn)

	<-ctx.Done()
	return nil
}

func sendEchoMessagePeriodically(ctx context.Context, conn *grpc.ClientConn) {
	client := pb.NewEchoClient(conn)
	tickerChan := time.NewTicker(5 * time.Second).C
	for {
		select {
		case t := <-tickerChan:
			message := fmt.Sprintf("Message at %d:%d:%d", t.Hour(), t.Minute(), t.Second())
			log.Printf("=================================================================")
			log.Printf("Sending %s", message)
			echoMessage(ctx, client, &pb.EchoRequest{Message: message})
		case <-ctx.Done():
			return
		}
	}
}

func echoMessage(ctx context.Context, client pb.EchoClient, echoRequest *pb.EchoRequest) {
	resp, err := client.EchoMessage(ctx, echoRequest)
	if err != nil {
		log.Fatalf("Could not send message : %v", err)
		return
	}
	log.Printf("Received %s from server %s", resp.Message, resp.ServerID)
}
