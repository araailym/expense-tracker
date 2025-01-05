package db

import (
	"database/sql"
	"fmt"
	"github.com/araailym/expense-tracker/internal/db/auth"
	"github.com/araailym/expense-tracker/internal/db/expense"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"strconv"
)

type DB struct {
	logger *slog.Logger
	pg     *sql.DB
	*expense.Expense
	*auth.Auth
}

func New(logger *slog.Logger) (*DB, error) {

	pgsql, err := NewPgSQL()
	if err != nil {
		return nil, err
	}

	return &DB{
		logger:  logger,
		pg:      pgsql,
		Expense: expense.New(pgsql, logger),
		Auth:    auth.New(pgsql, logger),
	}, nil
}

func NewPgSQL() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
