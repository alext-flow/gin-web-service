package mocks

import (
	"common/models"
)

var AlbumTypes map[string]models.AlbumType = map[string]models.AlbumType{
	"vinyl": {
		ID:     "vinyl-1",
		Name:  "vinyl",
		Desc:  "oldest and round",
	},
	"cd": {
		ID:     "cd-1",
		Name:  "cd",
		Desc:  "old, round and relatively small",
	},
	"minidisc": {
		ID:     "minidisc-1",
		Name:  "minidisc",
		Desc:  "round, small and the biggest flop of all",
	},
}
