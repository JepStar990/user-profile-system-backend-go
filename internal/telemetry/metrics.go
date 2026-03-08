package telemetry

import (
    "strconv"
    "time"

    "github.com/gofiber/adaptor/v2"
    "github.com/gofiber/fiber/v2"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

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

// InitMetrics registers Prometheus collectors.
// Call once at startup.
func InitMetrics() {
    // MustRegister will panic if called twice; if you run tests repeatedly,
    // prefer a sync.Once guard. For simplicity in a service binary, this is okay.
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// MetricsMiddleware instruments Fiber requests.
func MetricsMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()
        err := c.Next()

        method := c.Method()
        route := c.Route()
        path := ""
        if route != nil {
            path = route.Path
        }
        if path == "" {
            path = c.Path()
        }

        status := strconv.Itoa(c.Response().StatusCode())

        httpRequestsTotal.WithLabelValues(method, path, status).Inc()
        httpRequestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())

        return err
    }
}

// RegisterMetricsRoute mounts /metrics using promhttp via Fiber adaptor.
func RegisterMetricsRoute(app *fiber.App, path string) {
    app.Get(path, adaptor.HTTPHandler(promhttp.Handler()))
}
