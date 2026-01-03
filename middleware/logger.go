package middleware

import (
	"bytes"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
)

// responseWriter-обёртка для захвата статуса ответа
type statusRecorder struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b) // сохраняем копию ответа
	return rec.ResponseWriter.Write(b)
}

func RequestLogger(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// оборачиваем ResponseWriter, чтобы поймать статус
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK, body: bytes.NewBuffer(nil)}

			// вызываем следующий обработчик
			next.ServeHTTP(rec, r)

			duration := time.Since(start)
			reqID := middleware.GetReqID(r.Context())

			// аккуратный IP: берём первый из X-Forwarded-For, если есть
			clientIP := r.Header.Get("X-Forwarded-For")
			if clientIP != "" {
				clientIP = strings.Split(clientIP, ",")[0]
				clientIP = strings.TrimSpace(clientIP)
			} else {
				clientIP = r.RemoteAddr
			}

			// выбираем уровень логирования по статусу
			level := slog.LevelInfo
			if rec.status >= 500 {
				level = slog.LevelError
			} else if rec.status >= 400 {
				level = slog.LevelWarn
			} else {
				level = slog.LevelInfo // успешные запросы → INFO
			}

			// формируем лог
			logger := log
			if reqID != "" {
				logger = logger.With("request_id", reqID)
			}

			logger.Log(r.Context(), level, "HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rec.status,
				"remote", clientIP,
				"duration_ms", duration.Milliseconds(),
			)
		})
	}
}
