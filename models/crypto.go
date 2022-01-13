package models

type Crypto struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Votes int64
}
