package expenses

import (
	"fmt"
	"github.com/araailym/expense-tracker/internal/auth"
	"github.com/araailym/expense-tracker/internal/db/expense"
	"github.com/araailym/expense-tracker/pkg/httputils/request"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
)

type CreateExpenseRequest struct {
	Data *expense.ModelExpense `json:"data"`
}

type CreateExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) CreateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "CreateExpense")

	user, ok := ctx.Value("user").(*auth.UserData)
	if !ok {
		log.ErrorContext(
			ctx,
			"failed to type cast user data",
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("user: %+v\n", *user)

	requestBody := &CreateExpenseRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.InfoContext(ctx,
			"failed to parse request body",
			"error",
			err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.CreateExpense(ctx, requestBody.Data)
	if err != nil {
		log.InfoContext(ctx,
			"failed to query from db",
			"error",
			err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	if dbResp == nil {
		log.InfoContext(ctx,
			"row is empty",
		)
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}

	resp := CreateExpenseResponse{
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
		"success insert expense",
		"expense id",
		resp.Data.ID,
	)
	return
}
