package mocks

import (
	"common/models"
)

var Albums = []models.Album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
		// Type: models.AlbumType{
		// 	ID:   "vinyl-1",
		// 	Name: "vinyl",
		// 	Desc: "oldest and round",
		// },
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
		// Type: models.AlbumType{
		// 	ID:   "cd-1",
		// 	Name: "cd",
		// 	Desc: "old, round and relatively small",
		// },
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
		// Type: models.AlbumType{
		// 	ID:   "minidisc-1",
		// 	Name: "minidisc",
		// 	Desc: "round, small and the biggest flop of all",
		// },
	},
}