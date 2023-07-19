package models

type Album struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
	// Type   AlbumType `json:"type"`
}

type AlbumUpdate struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}