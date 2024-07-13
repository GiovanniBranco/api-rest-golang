package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GiovanniBranco/api-rest-golang/database/repositories"
	"github.com/GiovanniBranco/api-rest-golang/models"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := repositories.FindAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	movieId := getMovieIdFromRoute(r)
	movie := repositories.FindMovieById(movieId)
	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	movie = repositories.CreateMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieId := getMovieIdFromRoute(r)
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	repositories.UpdateMovie(movieId, movie)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieId := getMovieIdFromRoute(r)
	repositories.DeleteMovie(movieId)
	w.WriteHeader(http.StatusNoContent)
}

func getMovieIdFromRoute(r *http.Request) int {
	vars := mux.Vars(r)
	movieId, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	return movieId
}


