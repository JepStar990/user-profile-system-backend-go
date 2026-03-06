package http

import (
    "user-profile-system-backend-go/internal/controllers"
    "user-profile-system-backend-go/internal/security"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    // Controllers
    auth := controllers.AuthController{}
    profile := controllers.ProfileController{}
    settings := controllers.SettingsController{}

    api := app.Group("/api")

    //
    // PUBLIC AUTH ROUTES
    //
    authRoutes := api.Group("/auth")
    authRoutes.Post("/register", auth.Register)
    authRoutes.Post("/login", auth.Login)
    authRoutes.Post("/refresh", auth.Refresh)
    authRoutes.Post("/logout", auth.Logout)

    //
    // PRIVATE ROUTES (JWT REQUIRED)
    //
    private := api.Group("/private", security.AuthRequired)

    // Profile
    private.Get("/profile", profile.GetProfile)
    private.Put("/profile", profile.UpdateProfile)
    private.Post("/profile/change-password", profile.ChangePassword)
    private.Post("/profile/avatar", profile.UploadAvatar)

    // Settings
    private.Get("/settings", settings.GetSettings)
    private.Put("/settings/audio", settings.UpdateAudio)
    private.Put("/settings/voice", settings.UpdateVoice)
    private.Put("/settings/live", settings.UpdateLive)
    private.Put("/settings/notifications", settings.UpdateNotifications)
    private.Put("/settings/appearance", settings.UpdateAppearance)
    private.Put("/settings/privacy", settings.UpdatePrivacy)

    //
    // Health Check
    //
    api.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })
}
