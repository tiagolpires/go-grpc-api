package main

import (
	"context"
	"log"

	cryptoPb "go-grpc-api/proto/crypto"

	"go-grpc-api/config"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial(config.ServerPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer connection.Close()

	client := cryptoPb.NewCryptoServiceClient(connection)

	request := &cryptoPb.GetByIdRequest{Id: "1234"}

	response, err := client.GetCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not get crypto: %v", err)
	}

	log.Printf("Crypto: %s", response.Crypto)
}
