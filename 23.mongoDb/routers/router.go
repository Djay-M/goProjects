package routers

import (
	"github.com/Djay-M/mongoDB/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Status Apis
	router.HandleFunc("/", controllers.FetchServerStatus).Methods("GET")
	router.HandleFunc("/status", controllers.FetchServerStatus).Methods("GET")
	router.HandleFunc("/api/v1/status", controllers.FetchServerStatus).Methods("GET")

	// Movies Apis
	// V1
	router.HandleFunc("/api/v1/movies", controllers.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/v1/movies", controllers.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/v1/movies/{id}", controllers.MarkMovieAsWatchedController).Methods("PUT")
	router.HandleFunc("/api/v1/movies/byid/{id}", controllers.DeleteMovieByIdController).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/deleteall", controllers.DeleteAllMoviesController).Methods("DELETE")

	return router
}
