package expenses

import (
	"github.com/araailym/expense-tracker/internal/db/expense"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type FindExpensesResponse struct {
	Data []expense.ModelExpense `json:"data"`
}

func (h *Expenses) FindExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindExpenses")

	query := r.URL.Query()
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		log.ErrorContext(ctx, "fail parse query offset", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		log.ErrorContext(ctx, "fail parse query limit", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.FindExpenses(ctx, offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := FindExpensesResponse{
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
		"success find expenses",
		"number_of_expenses",
		len(resp.Data),
	)
	return
}
