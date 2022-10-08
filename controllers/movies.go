package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/iamcaleberic/semi-task/models"
	"github.com/iamcaleberic/semi-task/utils"
)

var movieCollectionName string = utils.CheckEnv("MONGODB_MOVIES_COLLECTION")
var movieCollection *mongo.Collection = models.GetCollection(movieCollectionName)

// GET /movies
func GetMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movies []*models.Movie

	cursor, err := movieCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result models.Movie
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, &result)
	}

	c.JSON(http.StatusOK, movies)

}

// GET /movies/:id
func GetMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movie models.Movie

	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	movieCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)

	c.JSON(http.StatusOK, movie)
}

// POST /movies
func CreateMovie(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.CreateMovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		Title:           input.Title,
		Genre:           input.Genre,
		Director:        input.Director,
		Language:        input.Language,
		CountryOfOrigin: input.CountryOfOrigin,
		Actors:          input.Actors,
		ReleaseDate:     input.ReleaseDate,
	}

	movieCollection.InsertOne(ctx, &movie)

	c.JSON(http.StatusOK, movie)
}
