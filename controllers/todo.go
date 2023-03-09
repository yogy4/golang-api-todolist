package controllers

import (
	"golang-api-todolist/entities"
	"golang-api-todolist/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	AddTodo(*gin.Context)
	GetAlltodos(*gin.Context)
	GetOnetodo(*gin.Context)
	UpdateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(s services.TodoService) TodoController {
	return todoController{
		todoService: s,
	}
}

// AddTodo implements TodoController
func (t todoController) AddTodo(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[TodoController]...Add todo")
	var todo entities.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while add todo"})
		return
	}
	todo, err := t.todoService.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while saving"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})
}

// DeleteTodo implements TodoController
func (t todoController) DeleteTodo(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[TodoController]...Delete todo")
	var todo entities.Todo
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	todo.ID = int(id)
	status := " Not Found"
	todo, err := t.todoService.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": status, "message": "Todo with ID " + ids + status})
}

// GetAlltodos implements TodoController
func (t todoController) GetAlltodos(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[TodoController]...Get all todos")
	// activitygroup := c.Params.ByName("activity_group_id")
	activitygroup := c.Query("activity_group_id")
	id, _ := strconv.Atoi(activitygroup)
	todos, err := t.todoService.GetAllTodos(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while getting todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

// GetOnetodo implements TodoController
func (t todoController) GetOnetodo(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[TodoController]... Get one todo")
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	todo, err := t.todoService.GetOneTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while getting todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})

}

// UpdateTodo implements TodoController
func (t todoController) UpdateTodo(c *gin.Context) {
	// panic("unimplemented")
	log.Print("[TodoController]...Update todo")
	var err error
	var todo entities.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	todo, err = t.todoService.UpdateTodo(id, todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while update todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})

}
