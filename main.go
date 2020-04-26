package main

import (
	"flag"
	"log"

	"github.com/zees-dev/go-grpc/internal/client"
	"github.com/zees-dev/go-grpc/internal/server"
)

func main() {
	typePtr := flag.String("type", "", "Denotes what type of grpc service to start: server, client")
	serverPortPtr := flag.Int("port", 50051, "Denotes gRPC server port")
	serverAddrPtr := flag.String("address", "localhost:50051", "Denotes gRPC server address in format: host:port")
	flag.Parse()

	switch *typePtr {
	case "server":
		server.RunServer(uint16(*serverPortPtr))
	case "client":
		client.RunClient(*serverAddrPtr, flag.Args()...)
	default:
		log.Printf("Unkown grpc service type: %q", *typePtr)
	}
}
