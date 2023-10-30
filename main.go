package main

import (
	"api-livro/src/repositories"
	"api-livro/src/routes"
)

func main() {
	repositories.ConnectionMongo()
	routes.HandleRequests()
}
