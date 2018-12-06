package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
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
	var err error
	if serverMode {
		log.Printf("server mode")
		err = NewServer(host).Run()
	} else {
		log.Printf("client mode")
		err = NewClient(host).Run()
	}

	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
