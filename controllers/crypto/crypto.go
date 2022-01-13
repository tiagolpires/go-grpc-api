package controller

import (
	"context"
	"go-grpc-api/database"
	"go-grpc-api/models"
	cryptoPb "go-grpc-api/proto/crypto"
)

type Server struct {
	cryptoPb.UnimplementedCryptoServiceServer
}

func (s *Server) UpvoteCryptoById(ctx context.Context, request *cryptoPb.CryptoIdRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	var crypto models.Crypto

	db.First(&crypto, request.Id)

	crypto.Votes = crypto.Votes + 1

	db.Save(&crypto)

	response := cryptoPb.Crypto{Id: crypto.ID, Code: crypto.Code, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}

func (s *Server) DownvoteCryptoById(ctx context.Context, request *cryptoPb.CryptoIdRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	var crypto models.Crypto

	db.First(&crypto, request.Id)

	crypto.Votes = crypto.Votes - 1

	db.Save(&crypto)

	response := cryptoPb.Crypto{Id: crypto.ID, Code: crypto.Code, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}

func (s *Server) GetCryptoById(ctx context.Context, request *cryptoPb.CryptoIdRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	var crypto models.Crypto

	db.First(&crypto, request.Id)

	response := cryptoPb.Crypto{Id: crypto.ID, Code: crypto.Code, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}

func (s *Server) ListCryptos(ctx context.Context, request *cryptoPb.EmptyRequest) (*cryptoPb.ListCryptosResponse, error) {
	db := database.GetDatabase()

	var cryptos []models.Crypto

	db.Find(&cryptos)

	var response []*cryptoPb.Crypto

	for _, crypto := range cryptos {
		cryptoCurrency := cryptoPb.Crypto{
			Id:    crypto.ID,
			Code:  crypto.Code,
			Name:  crypto.Name,
			Votes: crypto.Votes,
		}

		response = append(response, &cryptoCurrency)
	}

	return &cryptoPb.ListCryptosResponse{Cryptos: response}, nil
}

func (s *Server) CreateCrypto(ctx context.Context, request *cryptoPb.CreateCryptoRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	crypto := models.Crypto{Name: request.Name, Votes: request.Votes, Code: request.Code}

	db.Create(&crypto)

	response := cryptoPb.Crypto{Id: crypto.ID, Code: crypto.Code, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}

func (s *Server) UpdateCrypto(ctx context.Context, request *cryptoPb.UpdateCryptoRequest) (*cryptoPb.CryptoResponse, error) {
	db := database.GetDatabase()

	crypto := models.Crypto{ID: request.Id, Code: request.Code, Name: request.Name, Votes: request.Votes}

	db.Save(&crypto)

	response := cryptoPb.Crypto{Id: crypto.ID, Code: crypto.Code, Name: crypto.Name, Votes: crypto.Votes}

	return &cryptoPb.CryptoResponse{Crypto: &response}, nil
}

func (s *Server) DeleteCryptoById(ctx context.Context, request *cryptoPb.CryptoIdRequest) (*cryptoPb.SuccessMessageResponse, error) {
	db := database.GetDatabase()

	var crypto models.Crypto

	db.Delete(&crypto, request.Id)

	message := "Crypto was successfully deleted"

	return &cryptoPb.SuccessMessageResponse{Message: message, Success: true}, nil
}
