package controllers

import (
	"context"
	"net/http"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimmymuthoni/movies-stream/database"
	"github.com/jimmymuthoni/movies-stream/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/go-playground/validator/v10"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")
var validate = validator.New()

// function to get all movies
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		if movieCollection == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
			return
		}

		cursor, err := movieCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
			return
		}
		defer cursor.Close(ctx)

		var movies []models.Movie
		if err := cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}

//function to get one movie
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		movieID := c.Param("imdb_id")
		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required"})
			return
		}
		if movieCollection == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
			return
		}

		var movie models.Movie
		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			} else {
				log.Printf("Error fetching movie: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}

// function to add movie
func AddMovie() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
			return 
		}
		if err := validate.Struct(movie); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Validation failed","details": err.Error()})
			return 
		} 

		result, err := movieCollection.InsertOne(ctx, movie)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to add movie"})
			return 
		}
		c.JSON(http.StatusCreated, result)

	}
}