package expenses

import (
	"github.com/araailym/expense-tracker/internal/db/expense"
	"github.com/araailym/expense-tracker/pkg/httputils/request"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type UpdateExpenseRequest struct {
	Data *expense.ModelExpense `json:"data"`
}
type UpdateExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "UpdateExpense")

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

	requestBody := &UpdateExpenseRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.InfoContext(ctx,
			"failed to parse request body",
			"error",
			err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	if err := h.db.UpdateExpense(ctx, int64(id), requestBody.Data); err != nil {
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
		"success update expense",
		"id", id,
	)
	return
}
