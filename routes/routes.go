package routes

import (
	"log"
	"net/http"

	"github.com/GiovanniBranco/api-rest-golang/controllers"
	"github.com/GiovanniBranco/api-rest-golang/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.AddHeader)
	r.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", controllers.GetMovieById).Methods("GET")
	r.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
