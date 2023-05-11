package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"example.com/you/go-crud-api-test/pkg/models"
	"github.com/gorilla/mux"
)

type MovieHandler struct {
	Movie []models.Movie
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// encoding into json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// append all the other data in the place of the movie we want to delete
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// using _ for a variable we won't use
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// creating a random value from that range
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// this is not the way to do it when working with databases
	// but for this case it is a possible approach

	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(r)
	// loop over the movies, range
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie with the id youve sent
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// add a new movie - the movie that we send in the body of postman
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
