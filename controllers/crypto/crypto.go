package controller

import (
	"context"
	cryptoPb "go-grpc-api/proto/crypto"

	"log"
)

type Server struct {
	cryptoPb.UnimplementedCryptoServiceServer
}

func (s *Server) GetCryptoById(ctx context.Context, request *cryptoPb.GetByIdRequest) (*cryptoPb.CryptoResponse, error) {
	log.Printf("Received: %v", request.Id)
	crypto := cryptoPb.Crypto{Id: "id", Name: "Name", Votes: 1}
	return &cryptoPb.CryptoResponse{Crypto: &crypto}, nil
}
