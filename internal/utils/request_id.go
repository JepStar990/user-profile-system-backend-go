package utils

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// RequestID middleware ensures every request has a unique ID.
func RequestID() fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Get("X-Request-ID")
        if id == "" {
            id = uuid.New().String()
        }

        c.Locals("request_id", id)
        c.Set("X-Request-ID", id)

        return c.Next()
    }
}
