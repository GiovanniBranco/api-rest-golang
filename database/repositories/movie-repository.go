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

func CreateMovie(movie models.Movie) models.Movie {
	database.DB.Create(&movie)

	return movie
}

func DeleteMovie(id int) {
	database.DB.Delete(&models.Movie{}, id)
}

func UpdateMovie(id int, movie models.Movie) {
	movieDb := FindMovieById(id)

	movieDb.Title = movie.Title
	movieDb.Year = movie.Year

	database.DB.Save(&movieDb)
}
