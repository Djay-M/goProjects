package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Djay-M/mongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var counter = 0
var waitGroup sync.WaitGroup
var mux sync.Mutex

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func createMovie(movie models.Netflix) {
	createdMovie, createErr := models.Collections.InsertOne(context.Background(), bson.M{"movie": movie.Movie, "watched": movie.Watched})

	handleErrors(createErr)
	fmt.Println("Created a new movie", createdMovie)
}

func markMovieAsWatched(movieId string) int64 {
	id, idConversionErr := primitive.ObjectIDFromHex(movieId)

	handleErrors(idConversionErr)
	filter := bson.M{"_id": id}
	updateBson := bson.M{"$set": bson.M{"watched": true}}

	result, updateDbErr := models.Collections.UpdateOne(context.Background(), filter, updateBson)

	handleErrors(updateDbErr)
	fmt.Printf("Updated the movie with Id: %s as watched \n ", movieId)

	return result.MatchedCount
}

func deleteMovieById(movieId string) int {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteResult, deleteDbErr := models.Collections.DeleteOne(context.Background(), filter)

	handleErrors(deleteDbErr)
	return int(deleteResult.DeletedCount)
}

func deleteAllMovies() int {
	deleteResult, deleteErr := models.Collections.DeleteMany(context.Background(), bson.M{})
	handleErrors(deleteErr)

	return int(deleteResult.DeletedCount)
}

func fetchAllMovies() []primitive.M {
	cursor, findeErr := models.Collections.Find(context.Background(), bson.M{})
	handleErrors(findeErr)
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		decodeErr := cursor.Decode(&movie)
		handleErrors(decodeErr)

		movies = append(movies, movie)
	}

	return movies
}

// controller functions

func FetchServerStatus(response http.ResponseWriter, _ *http.Request) {
	counter += 1
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode("Server is up and running")
}

func GetAllMoviesController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movies := fetchAllMovies()

	json.NewEncoder(response).Encode(movies)
}

func CreateMovieController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newMovie models.Netflix

	json.NewDecoder(request.Body).Decode(&newMovie)
	createMovie(newMovie)

	json.NewEncoder(response).Encode("Created a new movie")
}

func MarkMovieAsWatchedController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	markMovieAsWatched(params["id"])

	json.NewEncoder(response).Encode("Movie Updated Successfully")
}

func DeleteMovieByIdController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	deleteMovieById(params["id"])

	json.NewEncoder(response).Encode("Delete a movie")
}

func DeleteAllMoviesController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	deleteAllMovies()

	json.NewEncoder(response).Encode("Deleted All movies from the DB")
}
