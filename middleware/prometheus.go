package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type statusRecorderPrometheus struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorderPrometheus) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations",
			Buckets: prometheus.DefBuckets, // стандартные buckets: 0.005s → 10s
		},
		[]string{"method", "path", "status"},
	)
)

func init() {
	// Регистрируем метрики
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// обёртка для статуса
		rec := &statusRecorderPrometheus{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()

		// обновляем метрики
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rec.status)).Inc()
		httpRequestDuration.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rec.status)).Observe(duration)
	})
}
