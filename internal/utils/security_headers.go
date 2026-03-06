package utils

import "github.com/gofiber/fiber/v2"

// SecurityHeaders sets best-practice security headers.
func SecurityHeaders() fiber.Handler {
    return func(c *fiber.Ctx) error {

        c.Set("X-Frame-Options", "DENY")
        c.Set("X-Content-Type-Options", "nosniff")
        c.Set("Referrer-Policy", "no-referrer")
        c.Set("X-XSS-Protection", "1; mode=block")
        c.Set("Content-Security-Policy", "default-src 'self'")

        return c.Next()
    }
}
