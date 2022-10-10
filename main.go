package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iamcaleberic/semi-task/controllers"
	"github.com/iamcaleberic/semi-task/models"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Movie Gallery API")
	})

	r.GET("/readyz", func(c *gin.Context) {
		models.Ping()
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.POST("/movies", controllers.CreateMovie)

	// listen and serve on 0.0.0.0:$PORT
	r.Run()
}
