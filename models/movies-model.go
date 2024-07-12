package models

type Movie struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var Movies []Movie
