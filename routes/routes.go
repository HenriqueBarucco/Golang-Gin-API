package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquebarucco/Golang-Gin-API/controllers"
)

func HandlerRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
