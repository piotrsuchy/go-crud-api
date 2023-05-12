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
	Movies []models.Movie
}

func (mh *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mh.Movies)
}

func (mh *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range mh.Movies {
		if item.ID == params["id"] {
			mh.Movies = append(mh.Movies[:index], mh.Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mh.Movies)
}

func (mh *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range mh.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func (mh *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	mh.Movies = append(mh.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func (mh *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range mh.Movies {
		if item.ID == params["id"] {
			mh.Movies = append(mh.Movies[:index], mh.Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			mh.Movies = append(mh.Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
