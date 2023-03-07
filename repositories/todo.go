package repositories

import (
	"golang-api-todolist/entities"
	"log"

	"gorm.io/gorm"
)

type todoRepository struct {
	DB *gorm.DB
}

type TodoRepository interface {
	CreateTodo(entities.Todo) (entities.Todo, error)
	GetAllTodos(int) ([]entities.Todo, error)
	GetOneTodo(int) (entities.Todo, error)
	UpdateTodo(int, entities.Todo) (entities.Todo, error)
	DeleteTodo(entities.Todo) (entities.Todo, error)
}

func NewTodoRepository(db *gorm.DB) todoRepository {
	return todoRepository{
		DB: db,
	}
}

func (t todoRepository) CreateTodo(todo entities.Todo) (entities.Todo, error) {
	log.Print("[TodoRepository]...CreateTodo")
	err := t.DB.Create(&todo).Error
	return todo, err
}

func (t todoRepository) GetAllTodos(id int) ([]entities.Todo, error) {
	log.Print("[TodoRepository]...GetAllTodos")
	var todos []entities.Todo
	err := t.DB.Find(&todos).Where("activity_group_id = ? ", id).Error
	return todos, err

}

func (t todoRepository) GetOneTodo(id int) (todo entities.Todo, err error) {
	log.Print("[TodoRepository]...GetOneTodo")
	err = t.DB.Where("todo_id = ?", id).First(&todo).Error
	return todo, err

}

func (t todoRepository) UpdateTodo(id int, todo entities.Todo) (entities.Todo, error) {
	log.Print("[TodoRepository]...UpdateTodo")
	var todosave entities.Todo
	err := t.DB.Where("todo_id = ?", id).First(&todosave).Error
	if err != nil {
		return todo, err
	}
	return todo, t.DB.Model(&todosave).Updates(todo).Error

}

func (t todoRepository) DeleteTodo(todo entities.Todo) (entities.Todo, error) {
	log.Print("[TodoRepository]...DeleteTodo")
	if err := t.DB.First(&todo, todo.ID).Error; err != nil {
		return todo, err
	}
	return todo, t.DB.Delete(&todo).Error
}
