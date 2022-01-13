package controller

import (
	"context"
	"go-grpc-api/database"
	"go-grpc-api/models"
	cryptoPb "go-grpc-api/proto/crypto"

	"log"
)

type Server struct {
	cryptoPb.UnimplementedCryptoServiceServer
}

func (s *Server) GetCryptoById(ctx context.Context, request *cryptoPb.GetByIdRequest) (*cryptoPb.CryptoResponse, error) {
	log.Printf("Received: %v", request.Id)
	crypto := cryptoPb.Crypto{Id: 1234, Name: "BNB", Votes: 1}
	return &cryptoPb.CryptoResponse{Crypto: &crypto}, nil
}

func (s *Server) CreateCrypto(ctx context.Context, request *cryptoPb.CreateCryptoRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	crypto := models.Crypto{Name: request.Name, Votes: request.Votes}

	db.Create(&crypto)

	response := cryptoPb.Crypto{Id: crypto.ID, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}
