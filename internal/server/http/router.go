package http

import (
    "user-profile-system-backend-go/internal/controllers"
    "user-profile-system-backend-go/internal/security"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    auth := controllers.AuthController{}
    profile := controllers.ProfileController{}
    settings := controllers.SettingsController{}
    favorites := controllers.FavoritesController{}
    downloads := controllers.DownloadsController{}

    api := app.Group("/api")

    // Auth
    authRoutes := api.Group("/auth")
    authRoutes.Post("/register", auth.Register)
    authRoutes.Post("/login", auth.Login)
    authRoutes.Post("/refresh", auth.Refresh)
    authRoutes.Post("/logout", auth.Logout)

    // Private
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

    // Favorites
    private.Post("/favorites", favorites.AddFavorite)
    private.Delete("/favorites", favorites.RemoveFavorite)
    private.Get("/favorites", favorites.ListFavorites)

    // Downloads
    private.Post("/downloads", downloads.RegisterDownload)
    private.Get("/downloads", downloads.ListDownloads)
    private.Delete("/downloads", downloads.RemoveDownload)
    private.Get("/downloads/url", downloads.PresignedURL)

    // Health
    api.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })
}
