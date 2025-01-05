package expenses

import (
	"github.com/araailym/expense-tracker/internal/db"
	"log/slog"
)

type Expenses struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Expenses {

	return &Expenses{
		logger: logger,
		db:     db,
	}

}
