package main

import (
	"log"
	"net"

	"go-grpc-api/config"
)

func main() {

	lis, err := net.Listen("tcp", config.ServerPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", lis.Addr())
	}

	log.Printf("Server listening at port %v", config.ServerPort)
}
