package main

import (
	"context"
	"fmt"
	"io"
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

	UpvoteCrypto(client)
	DownvoteCrypto(client)
	getCryptoStreamById(client)
	updateCrypto(client)
	getCryptoById(client)
	listCryptos(client)
	// createCrypto(client)
	// deleteCryptoById(client)
}

func UpvoteCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CryptoIdRequest{Id: 1}

	response, err := client.UpvoteCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not upvote crypto: %v", err)
	}

	fmt.Printf("UpvoteCrypto...\nCrypto: %s\n\n", response.Crypto)
}

func DownvoteCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CryptoIdRequest{Id: 1}

	response, err := client.DownvoteCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not downvote crypto: %v", err)
	}

	fmt.Printf("DownvoteCrypto...\nCrypto: %s\n\n", response.Crypto)
}

func getCryptoStreamById(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CryptoIdRequest{Id: 1}

	stream, err := client.GetCryptoStreamById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not get crypto stream: %v", err)
	}

	fmt.Println("Start crypto streaming...")

	for {
		crypto, err := stream.Recv()

		if err == io.EOF {
			fmt.Printf("Finish crypto streaming\n\n")
			break
		}

		if err != nil {
			log.Fatalf("an error ocurred while streaming crypto: %v", err)
		}
		fmt.Printf("%s Votes: %d\n", crypto.Crypto.Name, crypto.Crypto.Votes)
	}
}
func createCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CreateCryptoRequest{Name: "Devikins", Code: "DVK", Votes: 4}

	response, err := client.CreateCrypto(context.Background(), request)

	if err != nil {
		log.Fatalf("could not create crypto: %v", err)
	}

	fmt.Printf("CreateCrypto...\nCrypto: %s\n\n", response.Crypto)
}

func getCryptoById(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CryptoIdRequest{Id: 1}

	response, err := client.GetCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not get crypto: %v", err)
	}

	fmt.Printf("GetCryptoById...\nCrypto: %s\n\n", response.Crypto)
}

func updateCrypto(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.UpdateCryptoRequest{Id: 1, Code: "BTC", Name: "Bitcoin", Votes: 5}

	response, err := client.UpdateCrypto(context.Background(), request)

	if err != nil {
		log.Fatalf("could not update crypto: %v", err)
	}

	fmt.Printf("UpdateCrypto...\nCrypto: %s\n\n", response.Crypto)
}

func listCryptos(client cryptoPb.CryptoServiceClient) {
	empty := &cryptoPb.EmptyRequest{}
	response, err := client.ListCryptos(context.Background(), empty)

	if err != nil {
		log.Fatalf("could not list cryptos: %v", err)
	}

	fmt.Printf("ListCryptos...\nList: %s\n\n", response.Cryptos)
}

func deleteCryptoById(client cryptoPb.CryptoServiceClient) {
	request := &cryptoPb.CryptoIdRequest{Id: 1}

	response, err := client.DeleteCryptoById(context.Background(), request)

	if err != nil {
		log.Fatalf("could not delete crypto: %v", err)
	}

	fmt.Printf("DeleteCryptoById...\nResponse: %s\n\n", response)
}
