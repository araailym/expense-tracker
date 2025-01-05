package expenses

import (
	"github.com/araailym/expense-tracker/internal/db/expense"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type FindExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) FindExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindExpense")

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

	dbResp, err := h.db.FindExpense(ctx, int64(id))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail to query from db",
			"error", err,
		)
		http.Error(w, "fail to query from db", http.StatusInternalServerError)
		return
	}

	resp := FindExpenseResponse{
		Data: dbResp,
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
	); err != nil {
		log.ErrorContext(ctx, "fail json", "error", err)
		return
	}

	log.InfoContext(ctx,
		"success find expense",
		"expense id", resp.Data.ID,
	)
	return
}
