package repositories

import (
	"github.com/GiovanniBranco/api-rest-golang/database"
	"github.com/GiovanniBranco/api-rest-golang/models"
)

func FindAllMovies() []models.Movie {
	var movies []models.Movie

	database.DB.Find(&movies)

	return movies
}

func FindMovieById(id int) models.Movie {

	var movie models.Movie

	database.DB.First(&movie, id)

	return movie
}
