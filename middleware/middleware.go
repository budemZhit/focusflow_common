package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// --------------------
// Базовые middleware
// --------------------
func baseMiddlewares(log *slog.Logger) []func(http.Handler) http.Handler {
	// 1. Enrich — добавляем данные
	enrich := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
	}

	// 2. Observe — логирование
	observe := []func(http.Handler) http.Handler{
		RequestLogger(log),
	}

	// 3. Control — функциональные
	control := []func(http.Handler) http.Handler{
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
		CORSMiddleware,
	}

	// объединяем pipeline в правильном порядке
	var pipeline []func(http.Handler) http.Handler

	pipeline = append(pipeline, enrich...)
	pipeline = append(pipeline, observe...)
	pipeline = append(pipeline, control...)
	return pipeline
}

// --------------------
// Аутентификация
// --------------------
func authMiddlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		AuthMiddleware,
	}
}

// --------------------
// Observability pipeline (логи + метрики)
// --------------------
func pipelineMiddlewares(log *slog.Logger) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		PrometheusMiddleware,
	}
}

// --------------------
// Инициализация
// --------------------
func InitBase(r *chi.Mux, log *slog.Logger) *chi.Mux {
	mws := baseMiddlewares(log)
	for _, m := range mws {
		r.Use(m)
	}
	return r
}

func InitAuth(r *chi.Mux, log *slog.Logger) *chi.Mux {
	mws := append(baseMiddlewares(log), authMiddlewares()...)
	for _, m := range mws {
		r.Use(m)
	}
	return r
}

func InitPipeline(r *chi.Mux, log *slog.Logger) *chi.Mux {
	mws := append(baseMiddlewares(log), pipelineMiddlewares(log)...)
	for _, m := range mws {
		r.Use(m)
	}
	return r
}

func InitFull(r *chi.Mux, log *slog.Logger) *chi.Mux {
	mws := append(baseMiddlewares(log), pipelineMiddlewares(log)...)
	mws = append(mws, authMiddlewares()...)
	for _, m := range mws {
		r.Use(m)
	}
	return r
}
