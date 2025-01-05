package expense

import (
	"database/sql"
	"log/slog"
)

type Expense struct {
	logger *slog.Logger
	db     sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Expense {
	return &Expense{
		logger: logger,
		db:     *db,
	}
}
