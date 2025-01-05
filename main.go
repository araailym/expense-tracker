package main

import (
	"context"
	"fmt"
	"github.com/araailym/expense-tracker/internal/api"
	"github.com/araailym/expense-tracker/internal/db"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	_ = godotenv.Load()

	fmt.Println(os.Getenv("PORT"))

	d, err := db.New(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}
	//if err := d.Init(ctx); err != nil {
	//	panic(err)
	//}

	a := api.New(slog.With("service", "api"), d)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := a.Start(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to start api", "error", err.Error())
		}

		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	fmt.Println("shutting down gracefully")
}
