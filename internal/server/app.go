package server

import (
    "log"
    "os"
    "time"

    "user-profile-system-backend-go/internal/db"
    serverHttp "user-profile-system-backend-go/internal/server/http"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
    db.ConnectMySQL()
    utils.InitLogger()

    // ---- Telemetry init ----
    telemetry.InitMetrics()
    tracingShutdown, err := telemetry.InitTracing(context.Background(), "user-profile-system-backend-go")
    if err != nil {
        log.Fatalf("otel init error: %v", err)
    }

    app := fiber.New(fiber.Config{
        AppName:           "User Profile Backend",
        ErrorHandler:      utils.ErrorHandler,
        EnablePrintRoutes: true,
        ReadTimeout:       10 * time.Second,
        WriteTimeout:      10 * time.Second,
        IdleTimeout:       30 * time.Second,
    })

    // Panic recovery
    app.Use(recover.New())

    // Global middleware (request ID, security, logging, rate limit)
    serverHttp.RegisterGlobalMiddleware(app)

    // Telemetry middleware
    app.Use(telemetry.MetricsMiddleware())
    app.Use(telemetry.TracingMiddleware("user-profile-system-backend-go"))

    // CORS
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
        // AllowOrigins:     "*",
        AllowOrigins:     os.Getenv("CORS_ALLOW_ORIGINS"), // frontend link
        AllowMethods:     "GET,POST,PUT,DELETE",
        AllowHeaders:     "Authorization,Content-Type,X-Refresh-Token,X-Admin-Key,X-Request-ID",
    }))

    // Routes
    serverHttp.SetupRoutes(app)

    // Metrics endpoint (admin/internal)
    telemetry.RegisterMetricsRoute(app, "/metrics")

    // Graceful activity log shutdown
    go func() {
        <-app.Context().Done()
        services.ShutdownActivityLogger()
        _ = tracingShutdown(context.Background())
    }()

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("🚀 Server running on port:", port)
    log.Fatal(app.Listen(":" + port))
}
