package todoCRUD

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	todos := FindAllTodo()
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": todos})
}

func FindOne(c *gin.Context) {
	id := c.Param("id")
	todo := FindOneById(id)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": todo})
}

func Create(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CreateTodo(&todo)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "created"})
}

func Update(c *gin.Context) {
	var todoUpdate TodoUpdate
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTodo(id, todoUpdate)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "updated"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	DeleteTodo(id)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "deleted"})
}
