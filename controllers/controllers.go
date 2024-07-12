package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GiovanniBranco/api-rest-golang/database/repositories"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := repositories.FindAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	movie := repositories.FindMovieById(movieId)
	json.NewEncoder(w).Encode(movie)
}
