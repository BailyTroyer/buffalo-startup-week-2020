package util

import (
	"time"
	"net/http"
	"log"
	"strconv"

	"go.uber.org/zap"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// LoggingResponseWriter HTTP response log writer, storing HTTP status code from handlers
type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

// NewLoggingResponseWriter log response code from HTTP handlers
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{ResponseWriter: w}
}

// WriteHeader write HTTP status code header
func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

// InitLogger Initialize global zap logger
func InitLogger() {

	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zap.ReplaceGlobals(logger)

	defer logger.Sync()
}

// LoggingMiddleware structured logging middleware
func LoggingMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		loggingMiddlwareResponseWriter := newLoggingResponseWriter(w)
		h.ServeHTTP(loggingMiddlwareResponseWriter, r)
		zap.L().Info("Handled Request",
			zap.Int("durationMs", int(time.Now().Sub(startTime).Milliseconds())),
			zap.Int("responseCode", loggingMiddlwareResponseWriter.statusCode),
			zap.String("requestUrl", r.Host),
			zap.String("requestPath", r.URL.Path),
			zap.String("requestVerb", r.Method),
		)
    })
}

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})

	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of HTTP requests.",
	}, []string{"path", "method", "code"})
)

// PrometheusMiddleware implements mux.MiddlewareFunc.
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingMiddlwareResponseWriter := newLoggingResponseWriter(w)

		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(loggingMiddlwareResponseWriter, r)
		timer.ObserveDuration()
		httpRequestsTotal.WithLabelValues(path, r.Method, strconv.Itoa(loggingMiddlwareResponseWriter.statusCode)).Inc()
	})
}
