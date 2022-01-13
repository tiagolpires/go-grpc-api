package seeds

import (
	"go-grpc-api/models"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

var Seeds = []Seed{
	{
		Name: "CreateBTC",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Bitcoin", "BTC", 5)
		},
	},
	{
		Name: "CreateKLV",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Klever", "KLV", 6)
		},
	},
	{
		Name: "CreateDVK",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Devikins", "DVK", 3)
		},
	},
	{
		Name: "CreateETH",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Ethereum", "DVK", 10)
		},
	},
	{
		Name: "CreateBNB",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Binance", "BNB", 8)
		},
	},
	{
		Name: "CreateSAND",
		Run: func(db *gorm.DB) error {
			return CreateCrypto(db, "Sandbox", "SAND", 2)
		},
	},
}

func CreateCrypto(db *gorm.DB, name string, code string, votes int64) error {
	return db.Create(&models.Crypto{Name: name, Code: code, Votes: votes}).Error
}
