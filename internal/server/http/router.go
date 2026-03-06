package http

import (
    "user-profile-system-backend-go/internal/controllers"
    "user-profile-system-backend-go/internal/security"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    // Instantiate controllers
    auth := controllers.AuthController{}
    profile := controllers.ProfileController{}

    api := app.Group("/api")

    // ---------------------------
    // PUBLIC AUTH ROUTES
    // ---------------------------
    authRoutes := api.Group("/auth")
    authRoutes.Post("/register", auth.Register)
    authRoutes.Post("/login", auth.Login)
    authRoutes.Post("/refresh", auth.Refresh)
    authRoutes.Post("/logout", auth.Logout)

    // ---------------------------
    // PROTECTED ROUTES
    // ---------------------------
    private := api.Group("/private", security.AuthRequired)

    // PROFILE ROUTES
    private.Get("/profile", profile.GetProfile)
    private.Put("/profile", profile.UpdateProfile)
    private.Post("/profile/change-password", profile.ChangePassword)
    private.Post("/profile/avatar", profile.UploadAvatar)

    // HEALTH CHECK
    api.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })
}
