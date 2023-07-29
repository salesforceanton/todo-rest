package service

import (
	"github.com/go-exes/todo-serv"
	repository "github.com/go-exes/todo-serv/package/repositiory"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type TodoList interface {
	CreateList(input todo.TodoList, userId int) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId, id int) (todo.TodoList, error)
	DeleteList(id int) error
	UpdateList(input todo.UpdateListRequest, id int) (todo.TodoList, error)
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
