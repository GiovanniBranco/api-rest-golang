package main

import (
	"fmt"

	"github.com/GiovanniBranco/api-rest-golang/database"
	"github.com/GiovanniBranco/api-rest-golang/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Server is running on port 8080")

	routes.HandleRequest()

}
