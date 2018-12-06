package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

var (
	serverMode bool
	host       string
)

func init() {
	flag.BoolVar(&serverMode, "s", false, "run as the server")
	flag.StringVar(&host, "h", "0.0.0.0:8080", "the server's host")
	flag.Parse()
}

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Printf("shutdown signal received")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	var err error
	if serverMode {
		log.Printf("server mode")
		err = NewServer(host).Run(ctx)
	} else {
		log.Printf("client mode")
		err = NewClient(host).Run(ctx)
	}

	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
