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
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
    db.ConnectMySQL()
    utils.InitLogger()

    app := fiber.New(fiber.Config{
        AppName:               "User Profile Backend",
        EnablePrintRoutes:     true,
        ReadTimeout:           10 * time.Second,
        WriteTimeout:          10 * time.Second,
        IdleTimeout:           30 * time.Second,
        ErrorHandler:          utils.ErrorHandler,
    })

    // Middleware
    app.Use(recover.New())
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
        AllowOrigins:     "*",
        AllowMethods:     "GET,POST,PUT,DELETE",
        AllowHeaders:     "Authorization,Content-Type,X-Refresh-Token",
    }))

    // Register routes
    serverHttp.SetupRoutes(app)

    // Graceful shutdown for async logging worker
    go func() {
        <-app.Context().Done()
        services.ShutdownActivityLogger()
    }()

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("🚀 Server running on port:", port)
    log.Fatal(app.Listen(":" + port))
}
