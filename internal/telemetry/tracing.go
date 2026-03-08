package telemetry

import (
    "context"
    "os"
    "time"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// InitTracing sets up OpenTelemetry tracing.
// Exports OTLP traces to a collector (e.g. Grafana Tempo, OpenTelemetry Collector).
// You control endpoint via OTEL_EXPORTER_OTLP_ENDPOINT env var.
// Returns shutdown function to call on graceful shutdown.
func InitTracing(ctx context.Context, serviceName string) (func(context.Context) error, error) {
    endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
    if endpoint == "" {
        // Tracing disabled if not configured.
        otel.SetTracerProvider(sdktrace.NewTracerProvider())
        return func(context.Context) error { return nil }, nil
    }

    // OTLP HTTP exporter
    exp, err := otlptracehttp.New(ctx,
        otlptracehttp.WithEndpoint(endpoint),
        otlptracehttp.WithInsecure(), // set to secure + TLS in production when you have certs
    )
    if err != nil {
        return nil, err
    }

    res, err := resource.New(ctx,
        resource.WithAttributes(
            semconv.ServiceName(serviceName),
            semconv.DeploymentEnvironment(os.Getenv("APP_ENV")),
        ),
    )
    if err != nil {
        return nil, err
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp,
            sdktrace.WithBatchTimeout(2*time.Second),
        ),
        sdktrace.WithResource(res),
    )

    otel.SetTracerProvider(tp)
    return tp.Shutdown, nil
}
