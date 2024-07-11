package main

import (
	"fmt"

	"github.com/GiovanniBranco/api-rest-golang/models"
	"github.com/GiovanniBranco/api-rest-golang/routes"
)

func main() {
	fmt.Println("Server is running on port 8080")
	models.Movies = []models.Movie{
		{
			Title: "The Shawshank Redemption",
			Year:  1994,
		},
		{
			Title: "The Godfather",
			Year:  1972,
		},
	}
	routes.HandleRequest()

}
