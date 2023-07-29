package repository

import (
	"fmt"

	"github.com/go-exes/todo-serv"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) GetAllLists(userId int) ([]todo.TodoList, error) {
	var result []todo.TodoList

	query := fmt.Sprintf("SELECT list.id, list.title, list.description FROM %s list INNER JOIN %s userlist on list.id = userlist.list_id WHERE userlist.user_id = $1", todoListsTable, userListsTable)
	err := r.db.Select(&result, query, userId)

	return result, err
}

func (r *TodoListPostgres) CreateList(input todo.TodoList, userId int) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// Create Todo List record
	var listId int

	createListRecordQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := transaction.QueryRow(createListRecordQuery, input.Title, input.Description)
	if err := row.Scan(&listId); err != nil {
		transaction.Rollback()
		return 0, err
	}

	// Create user list junction record
	createUserListRecordQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	_, err = transaction.Exec(createUserListRecordQuery, userId, listId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return listId, transaction.Commit()
}

func (r *TodoListPostgres) GetListById(userId, id int) (todo.TodoList, error) {
	var result todo.TodoList

	query := fmt.Sprintf(
		`SELECT list.id, list.title, list.description 
		FROM %s list INNER JOIN %s userlist on list.id = userlist.list_id 
		WHERE userlist.user_id = $1 AND list.id = $2`,
		todoListsTable, userListsTable,
	)
	err := r.db.Get(&result, query, userId, id)

	return result, err
}

func (r *TodoListPostgres) DeleteList(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s list WHERE list.id = $1`, todoListsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *TodoListPostgres) UpdateList(input todo.UpdateListRequest, id int) (todo.TodoList, error) {
	var result todo.TodoList

	query := fmt.Sprintf(
		`UPDATE %s SET title='%s', description='%s' WHERE id='%d' RETURNING id, title, description`,
		todoListsTable, input.Title, input.Description, id,
	)
	fmt.Println(query)
	err := r.db.Get(&result, query)

	return result, err
}
