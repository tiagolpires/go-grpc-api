package controller

import (
	"context"
	"go-grpc-api/database"
	"go-grpc-api/models"
	"testing"

	cryptoPb "go-grpc-api/proto/crypto"

	"gorm.io/gorm"
)

func TestUpvoteCryptoById(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.CryptoIdRequest{Id: newCrypto.ID}

	response, err := s.UpvoteCryptoById(context.Background(), request)

	if err != nil {
		t.Fatal("error calling UpvoteCryptoById")
	}

	if response.Crypto.Votes != newCrypto.Votes+1 {
		t.Fatal("error: UpvoteCryptoById did not upvote")
	}
}

func TestDownvoteCryptoById(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.CryptoIdRequest{Id: newCrypto.ID}

	response, err := s.DownvoteCryptoById(context.Background(), request)

	if err != nil {
		t.Fatal("error calling DownvoteCryptoById")
	}

	if response.Crypto.Votes != newCrypto.Votes-1 {
		t.Fatal("error: DownvoteCryptoById did not downvote")
	}
}

func TestGetCryptoById(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.CryptoIdRequest{Id: newCrypto.ID}

	response, err := s.GetCryptoById(context.Background(), request)

	if err != nil {
		t.Fatal("error calling GetCryptoById")
	}

	if response.Crypto.Id != newCrypto.ID {
		t.Fatal("error: GetCryptoById returns incorrect crypto")
	}
}

func TestListCryptos(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.EmptyRequest{}

	response, err := s.ListCryptos(context.Background(), request)

	if err != nil {
		t.Fatal("error calling ListCryptos")
	}

	if len(response.Cryptos) <= 0 {
		t.Fatal("error: ListCryptos returns a list with incorrect length")
	}
}

func TestCreateCrypto(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	crypto := cryptoPb.CreateCryptoRequest{Name: "Test", Votes: 2, Code: "TEST"}

	response, err := s.CreateCrypto(context.Background(), &crypto)
	defer db.Delete(&models.Crypto{}, response.Crypto.Id)

	if err != nil {
		t.Fatal("error calling CreateCrypto")
	}

	if response.Crypto.Code != "TEST" {
		t.Fatal("error: CreateCrypto returns incorrect crypto value")
	}
}

func TestUpdateCrypto(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.UpdateCryptoRequest{Id: newCrypto.ID, Name: "Test", Code: "UPDT"}

	response, err := s.UpdateCrypto(context.Background(), request)

	if err != nil {
		t.Fatal("error calling UpdateCrypto")
	}

	if response.Crypto.Code != request.Code {
		t.Fatal("error: UpdateCrypto returns incorrect value")
	}
}

func TestDeleteCryptoById(t *testing.T) {
	s := Server{}
	database.StartDB()
	defer database.CloseConn()
	db := database.GetDatabase()

	newCrypto := createCrypto(db)
	defer db.Delete(&newCrypto)

	request := &cryptoPb.CryptoIdRequest{Id: newCrypto.ID}

	response, err := s.DeleteCryptoById(context.Background(), request)

	if err != nil {
		t.Fatal("error calling DeleteCryptoById")
	}

	if !response.Success {
		t.Fatal("error: DeleteCryptoById return success false")
	}
}

func createCrypto(db *gorm.DB) *models.Crypto {

	crypto := models.Crypto{Name: "Test", Votes: 2, Code: "TEST"}

	db.Create(&crypto)

	return &crypto
}
