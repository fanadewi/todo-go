package main

import (
	"net/http"

	"github.com/fanadewi/todo-go/controllers"
	"github.com/fanadewi/todo-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"reply": "pong"})
	})
	r.GET("/todos/:id", controllers.FindTodo)
	r.GET("/todos", controllers.FindTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.Run()
}
