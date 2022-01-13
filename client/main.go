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

	createCrypto(client)
	getCryptoById(client)
}

func createCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CreateCryptoRequest{Name: "Devikins", Votes: 4}

	response, err := client.CreateCrypto(context.Background(), request)

	if err != nil {
		log.Fatalf("could not create crypto: %v", err)
	}

	log.Printf("CreateCrypto output: %s", response.Crypto)
}

func getCryptoById(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.GetByIdRequest{Id: 1}

	response, err := client.GetCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not get crypto: %v", err)
	}

	log.Printf("GetCryptoById output: %s", response.Crypto)
}
