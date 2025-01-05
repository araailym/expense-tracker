package handler

import (
	"github.com/araailym/expense-tracker/internal/api/handler/auth"
	"github.com/araailym/expense-tracker/internal/api/handler/expenses"
	"github.com/araailym/expense-tracker/internal/db"
	"log/slog"
)

type Handler struct {
	*auth.Auth
	*expenses.Expenses
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Auth:     auth.New(logger, db),
		Expenses: expenses.New(logger, db),
	}
}
