package models

import "gorm.io/gorm"

// album represents data about a record album.

type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Check  string  `json:"check"`
}

func (album Album) Init() []Album {
	// albums slice to seed record album data.
	albums := []Album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return albums
}
