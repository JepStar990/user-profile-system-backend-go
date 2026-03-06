package http

import (
    "user-profile-system-backend-go/internal/controllers"
    "user-profile-system-backend-go/internal/security"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    auth := controllers.AuthController{}

    api := app.Group("/api")

    api.Post("/auth/register", auth.Register)
    api.Post("/auth/login", auth.Login)
    api.Post("/auth/refresh", auth.Refresh)
    api.Post("/auth/logout", auth.Logout)

    private := api.Group("/private", security.AuthRequired)
    private.Get("/profile", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"user_id": c.Locals("user_id")})
    })
}
