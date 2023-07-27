package models

import (
	"github.com/google/uuid"
)

type Album struct {
	ID     string    	`gorm:"type:char(36);primaryKey" json:"id"`
	Title  string    	`json:"title"`
	Artist string    	`json:"artist"`
	Price  float64   	`json:"price"`
	AlbumType AlbumType `gorm:"foreignKey:TypeID" json:"type"`
	TypeID uuid.UUID 	`gorm:"type:char(36)" json:"-"`
}

type AlbumUpdate struct {
	ID     string   `json:"id,omitempty"`
	Title  string   `json:"title,omitempty"`
	Artist string   `json:"artist,omitempty"`
	Price  float64  `json:"price,omitempty"`
}