package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jimmymuthoni/movies-stream/database"
	"github.com/jimmymuthoni/movies-stream/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies(client *mongo.Client) gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()	
		cursor, err := movieCollection.Find(ctx, bson.D{})
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		}
		defer cursor.Close(ctx)

		var movies []models.Movie
		if err = cursor.All(ctx, &movies); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
			return 
		}

		c.JSON(http.StatusOK, movies)

	}
}