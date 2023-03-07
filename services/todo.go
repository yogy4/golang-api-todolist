package services

import (
	"golang-api-todolist/entities"
	"golang-api-todolist/repositories"
)

type todoService struct {
	todoRepository repositories.TodoRepository
}

type TodoService interface {
	CreateTodo(entities.Todo) (entities.Todo, error)
	GetAllTodos(int) ([]entities.Todo, error)
	GetOneTodo(int) (entities.Todo, error)
	UpdateTodo(int, entities.Todo) (entities.Todo, error)
	DeleteTodo(entities.Todo) (entities.Todo, error)
}

func NewTodoService(r repositories.TodoRepository) TodoService {
	return todoService{
		todoRepository: r,
	}
}

// CreateTodo implements TodoService
func (t todoService) CreateTodo(todo entities.Todo) (entities.Todo, error) {
	// panic("unimplemented")
	return t.todoRepository.CreateTodo(todo)
}

// DeleteTodo implements TodoService
func (t todoService) DeleteTodo(todo entities.Todo) (entities.Todo, error) {
	// panic("unimplemented")
	return t.todoRepository.DeleteTodo(todo)
}

// GetAllTodos implements TodoService
func (t todoService) GetAllTodos(activitygroup int) ([]entities.Todo, error) {
	// panic("unimplemented")
	return t.todoRepository.GetAllTodos(activitygroup)
}

// GetOneTodo implements TodoService
func (t todoService) GetOneTodo(id int) (entities.Todo, error) {
	// panic("unimplemented")
	return t.todoRepository.GetOneTodo(id)
}

// UpdateTodo implements TodoService
func (t todoService) UpdateTodo(id int, todo entities.Todo) (entities.Todo, error) {
	// panic("unimplemented")
	return t.todoRepository.UpdateTodo(id, todo)
}
