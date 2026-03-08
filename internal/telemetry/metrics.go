package telemetry

import (
    "strconv"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// We expose two types of metrics:
// 1) HTTP request count
// 2) HTTP request duration histogram
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests processed.",
        },
        []string{"method", "path", "status"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Histogram of HTTP request durations (seconds).",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "path"},
    )
)

// InitMetrics registers collectors globally.
// Call once at startup.
func InitMetrics() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// MetricsMiddleware instruments Fiber requests.
// IMPORTANT: Avoid high-cardinality labels. We use c.Route().Path (template path) not raw URL.
func MetricsMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()
        err := c.Next()

        method := c.Method()
        path := c.Route().Path
        status := strconv.Itoa(c.Response().StatusCode())

        httpRequestsTotal.WithLabelValues(method, path, status).Inc()
        httpRequestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())

        return err
    }
}

// RegisterMetricsRoute mounts /metrics using promhttp.
func RegisterMetricsRoute(app *fiber.App, path string) {
    // Fiber doesn't natively accept http.Handler, so we use a bridge.
    app.Get(path, func(c *fiber.Ctx) error {
        promhttp.Handler().ServeHTTP(c.Context().ResponseWriter(), c.Context().Request())
        return nil
    })
}
