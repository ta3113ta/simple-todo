package main

import (
	"todo/todoCRUD"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/todo", todoCRUD.FindAll)
	router.GET("/todo/:id", todoCRUD.FindOne)
	router.POST("/todo", todoCRUD.Create)
	router.PATCH("/todo/:id", todoCRUD.Update)
	router.DELETE("/todo/:id", todoCRUD.Delete)

	router.Run(":5000")
}
