package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/cufxt/ai-certificate-portfolio/backend/gen/health/v1/healthv1connect"
	"github.com/cufxt/ai-certificate-portfolio/backend/internal/config"
	"github.com/cufxt/ai-certificate-portfolio/backend/internal/handler"
)

func main() {
	setupLogger()

	cfg := config.Load()
	router := setupRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	runServer(srv, cfg)
	waitForShutdown(srv)
}

func setupLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	healthHandler := handler.NewHealthHandler()
	path, connectHandler := healthv1connect.NewHealthServiceHandler(healthHandler)
	r.Mount(path, connectHandler)

	return r
}

func runServer(srv *http.Server, cfg config.Config) {
	go func() {
		slog.Info("starting server", "port", cfg.Port, "env", cfg.Env)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()
}

func waitForShutdown(srv *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	slog.Info("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
	}
}
