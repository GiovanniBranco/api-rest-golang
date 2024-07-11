package routes

import (
	"log"
	"net/http"

	"github.com/GiovanniBranco/api-rest-golang/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/movies", controllers.GetAllMovies)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))

}
