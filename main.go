package main

import (
	"go-grpc-api/config"
	cryptoController "go-grpc-api/controllers/crypto"
	cryptoPb "go-grpc-api/proto/crypto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", config.ServerPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cryptoPb.RegisterCryptoServiceServer(grpcServer, &cryptoController.Server{})

	log.Printf("Server listening at port %v", config.ServerPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
