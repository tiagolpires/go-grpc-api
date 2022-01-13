package migrations

import (
	"go-grpc-api/models"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.Crypto{})
}
