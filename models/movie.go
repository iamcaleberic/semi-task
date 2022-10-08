package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title           string             `json:"title"`
	Genre           string             `json:"genre"`
	Director        string             `json:"director"`
	Language        string             `json:"language"`
	CountryOfOrigin string             `json:"country_of_origin"`
	Actors          []string           `json:"actors"`
	ReleaseDate     string             `json:"release_date"`
}

type CreateMovieInput struct {
	Title           string   `json:"title" binding:"required"`
	Genre           string   `json:"genre" binding:"required"`
	Director        string   `json:"director" binding:"required"`
	Language        string   `json:"language" binding:"required"`
	CountryOfOrigin string   `json:"country_of_origin" binding:"required"`
	Actors          []string `json:"actors" binding:"required"`
	ReleaseDate     string   `json:"release_date" binding:"required"`
}
