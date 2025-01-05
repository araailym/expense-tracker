package expenses

import (
	"github.com/araailym/expense-tracker/internal/db/expense"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type DeleteExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "DeleteExpense")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail to convert id",
			"error", err,
		)
		http.Error(w, "fail to convert id", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteExpense(ctx, int64(id)); err != nil {
		log.ErrorContext(
			ctx,
			"fail to query from db",
			"error", err,
		)
		http.Error(w, "fail to query from db", http.StatusInternalServerError)
		return
	}

	if err := response.JSON(
		w,
		http.StatusNoContent,
		nil,
	); err != nil {
		log.ErrorContext(ctx, "fail json", "error", err)
		return
	}

	log.InfoContext(ctx,
		"success delete expense",
		"id", id,
	)
	return
}
