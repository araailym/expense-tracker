package router

import (
	"context"
	"net/http"
)

func (r *Router) expenses(ctx context.Context) {
	r.router.Handle("GET /expenses", http.HandlerFunc(r.handler.FindExpenses))
	r.router.Handle("GET /expenses/{id}", http.HandlerFunc(r.handler.FindExpense))
	r.router.Handle("POST /expenses", http.HandlerFunc(r.handler.CreateExpense))
	r.router.Handle("PUT /expenses/{id}", http.HandlerFunc(r.handler.UpdateExpense))
	r.router.Handle("DELETE /expenses/{id}", http.HandlerFunc(r.handler.DeleteExpense))
}
