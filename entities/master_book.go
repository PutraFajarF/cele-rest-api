package entities

import (
	"gorm.io/gorm"
)

type MasterBook struct {
	gorm.Model
	AuthorID uint   `json:"author_id" gorm:"foreignKey:ID"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Price    int    `json:"price"`
	// BookTransaction []BookTransaction `gorm:"foreignKey:BookID;references:ID"``
}
