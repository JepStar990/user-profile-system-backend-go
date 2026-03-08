package telemetry

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

func TracingMiddleware(serviceName string) fiber.Handler {
    tracer := otel.Tracer(serviceName)

    return func(c *fiber.Ctx) error {
        routePath := c.Route().Path
        ctx := context.Background()

        ctx, span := tracer.Start(ctx, routePath)
        span.SetAttributes(
            attribute.String("http.method", c.Method()),
            attribute.String("http.route", routePath),
            attribute.String("http.client_ip", c.IP()),
            attribute.String("http.user_agent", string(c.Context().UserAgent())),
        )

        err := c.Next()

        span.SetAttributes(attribute.Int("http.status_code", c.Response().StatusCode()))
        if err != nil {
            span.SetAttributes(attribute.String("error", err.Error()))
        }
        span.End()

        return err
    }
}
