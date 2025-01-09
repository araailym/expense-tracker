package router

import (
	"context"
	"net/http"
)

func (r *Router) expenses(ctx context.Context) {
	r.router.Handle(
		"GET /expenses",
		r.midd.Authenticator(http.HandlerFunc(r.handler.FindExpenses)),
	)
	r.router.Handle(
		"GET /expenses/{id}",
		r.midd.Authenticator(http.HandlerFunc(r.handler.FindExpense)),
	)
	r.router.Handle(
		"POST /expenses",
		r.midd.Authorizer("admin", http.HandlerFunc(r.handler.CreateExpense)),
	)
	r.router.Handle(
		"PUT /expenses/{id}",
		r.midd.Authorizer("admin", http.HandlerFunc(r.handler.UpdateExpense)))
	r.router.Handle(
		"DELETE /expenses/{id}",
		r.midd.Authorizer("admin", http.HandlerFunc(r.handler.DeleteExpense)),
	)

}
