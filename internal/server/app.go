package server

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "user-profile-system-backend-go/internal/db"
    serverHttp "user-profile-system-backend-go/internal/server/http"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/telemetry"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
    db.ConnectMySQL()

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

    // CORS – allow specific origins from env, or reflect any origin for dev
    corsOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
    var corsConfig cors.Config
    if corsOrigins == "" {
        corsConfig = cors.Config{
            AllowCredentials: true,
            AllowOriginsFunc: func(origin string) bool { return true },
            AllowMethods:     "GET,POST,PUT,DELETE",
            AllowHeaders:     "Authorization,Content-Type,X-Refresh-Token,X-Admin-Key,X-Request-ID",
        }
    } else {
        corsConfig = cors.Config{
            AllowCredentials: true,
            AllowOrigins:     corsOrigins,
            AllowMethods:     "GET,POST,PUT,DELETE",
            AllowHeaders:     "Authorization,Content-Type,X-Refresh-Token,X-Admin-Key,X-Request-ID",
        }
    }
    app.Use(cors.New(corsConfig))

    // Routes
    serverHttp.SetupRoutes(app)

    // Metrics endpoint (admin/internal)
    telemetry.RegisterMetricsRoute(app, "/metrics")

    // Graceful shutdown on SIGINT/SIGTERM
    go func() {
        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
        <-sigCh
        log.Println("Shutting down...")
        services.ShutdownActivityLogger()
        _ = tracingShutdown(context.Background())
        _ = app.Shutdown()
    }()

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("🚀 Server running on port:", port)
    log.Fatal(app.Listen(":" + port))
}
