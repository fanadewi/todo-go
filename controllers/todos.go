package controllers

import (
	"net/http"

	"github.com/fanadewi/todo-go/models"
	"github.com/gin-gonic/gin"
)

// CreateTodoInput for validation
type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

// UpdateTodoInput for validation
type UpdateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

// FindTodos ...
func FindTodos(c *gin.Context) {
	var todos []models.Todo
	// if both params filled
	states, ok := c.Request.URL.Query()["state"]
	titles, ok2 := c.Request.URL.Query()["title"]
	statePresent := ok && len(states[0]) > 0
	titlePresent := ok2 && len(titles[0]) > 0

	if statePresent && titlePresent {
		models.DB.Where("state = ? and title LIKE ?", states[0], titles[0]).Find(&todos)
	} else if statePresent {
		models.DB.Where("state = ?", states[0]).Find(&todos)
	} else if titlePresent {
		models.DB.Where("title LIKE ?", titles[0]).Find(&todos)
	} else {
		models.DB.Find(&todos)
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// CreateTodo ...
func CreateTodo(c *gin.Context) {
	// Validate input
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Create todo
	todo := models.Todo{Title: input.Title}
	todo.State = "todo"
	models.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

// FindTodo ...
func FindTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// UpdateTodo ...
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Model(&todo).Update("title", input.Title)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DeleteTodo ...
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)

	c.Writer.WriteHeader(http.StatusOK)
}
