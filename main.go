package main

import (
	"github.com/henriquebarucco/Golang-Gin-API/database"
	"github.com/henriquebarucco/Golang-Gin-API/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandlerRequests()
}
