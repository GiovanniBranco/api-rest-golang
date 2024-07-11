package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GiovanniBranco/api-rest-golang/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Movies)

}
