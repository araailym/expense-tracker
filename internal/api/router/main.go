package router

import (
	"context"
	"github.com/araailym/expense-tracker/internal/api/handler"
	"github.com/araailym/expense-tracker/internal/api/middleware"
	"net/http"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(handler *handler.Handler, midd *middleware.Middleware) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
		midd:    midd,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.auth(ctx)
	r.expenses(ctx)

	return r.router
}
