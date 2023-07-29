package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	todoItemsTable  = "todo_items"
	userListsTable  = "users_lists"
	itemsListsTable = "items_lists"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	pgUrl, err := pq.ParseURL(fmt.Sprintf("postgres://%s:%s@stampy.db.elephantsql.com/%s", cfg.Username, cfg.Password, cfg.DBName))

	db, err := sqlx.Open("postgres", pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
