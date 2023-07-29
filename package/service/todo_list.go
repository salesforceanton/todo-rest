package service

import (
	"github.com/go-exes/todo-serv"
	repository "github.com/go-exes/todo-serv/package/repositiory"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) CreateList(input todo.TodoList, userId int) (int, error) {
	return s.repo.CreateList(input, userId)
}

func (s *TodoListService) GetListById(userId, id int) (todo.TodoList, error) {
	return s.repo.GetListById(userId, id)
}

func (s *TodoListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}

func (s *TodoListService) UpdateList(input todo.UpdateListRequest, id int) (todo.TodoList, error) {
	return s.repo.UpdateList(input, id)
}
