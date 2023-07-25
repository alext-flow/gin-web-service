package models
import (
	"github.com/google/uuid"
)


type AlbumType struct {
	ID   uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}