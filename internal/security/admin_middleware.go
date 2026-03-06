package security

import (
    "os"

    "github.com/gofiber/fiber/v2"
)

// AdminAuthRequired ensures only admins (internal systems) can access admin endpoints.
func AdminAuthRequired(c *fiber.Ctx) error {
    apiKey := c.Get("X-Admin-Key")
    if apiKey == "" {
        return fiber.NewError(fiber.StatusUnauthorized, "missing admin key")
    }

    expected := os.Getenv("ADMIN_API_KEY")
    if expected == "" {
        return fiber.NewError(fiber.StatusInternalServerError, "admin key not configured")
    }

    if apiKey != expected {
        return fiber.NewError(fiber.StatusForbidden, "invalid admin key")
    }

    return c.Next()
}
