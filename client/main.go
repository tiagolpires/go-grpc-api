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
	updateCrypto(client)
	listCryptos(client)
	deleteCryptoById(client)
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
	request := &cryptoPb.GetCryptoByIdRequest{Id: 1}

	response, err := client.GetCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not get crypto: %v", err)
	}

	log.Printf("GetCryptoById output: %s", response.Crypto)
}

func updateCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.UpdateCryptoRequest{Id: 1, Name: "Bitcoin", Votes: 5}

	response, err := client.UpdateCrypto(context.Background(), request)

	if err != nil {
		log.Fatalf("could not update crypto: %v", err)
	}

	log.Printf("UpdateCrypto output: %s", response.Crypto)
}

func listCryptos(client cryptoPb.CryptoServiceClient) {
	empty := &cryptoPb.EmptyRequest{}
	response, err := client.ListCryptos(context.Background(), empty)

	if err != nil {
		log.Fatalf("could not list cryptos: %v", err)
	}

	log.Printf("ListCryptos output: %s", response.Cryptos)
}

func deleteCryptoById(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.GetCryptoByIdRequest{Id: 1}

	response, err := client.DeleteCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not delete crypto: %v", err)
	}

	log.Printf("DeleteCryptoById output: %s", response)
}
