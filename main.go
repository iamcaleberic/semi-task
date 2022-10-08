package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iamcaleberic/semi-task/controllers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Movie Gallery API")
	})

	r.GET("/movies", controllers.GetMovies)
	r.POST("/movies", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.GetMovie)

	// listen and serve on 0.0.0.0:$PORT
	r.Run()
}
