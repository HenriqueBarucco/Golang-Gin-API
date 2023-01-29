package main

import (
	"github.com/henriquebarucco/Golang-Gin-API/database"
	"github.com/henriquebarucco/Golang-Gin-API/models"
	"github.com/henriquebarucco/Golang-Gin-API/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Henrique", CPF: "00000000001", RG: "1000000000"},
		{Nome: "Teste", CPF: "00000000002", RG: "2000000000"},
		{Nome: "Barucco", CPF: "00000000003", RG: "3000000000"},
	}
	routes.HandlerRequests()
}
