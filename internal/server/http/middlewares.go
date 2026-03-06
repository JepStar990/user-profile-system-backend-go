package http

import (
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
)

// RegisterGlobalMiddleware attaches all global layers to the app.
func RegisterGlobalMiddleware(app *fiber.App) {

    // Assign request ID first
    app.Use(utils.RequestID())

    // Panic safety
    app.Use(utils.SecurityHeaders())

    // Structured request logging
    app.Use(utils.RequestLogger())

    // Global rate limiter
    app.Use(utils.RateLimiter())
}
