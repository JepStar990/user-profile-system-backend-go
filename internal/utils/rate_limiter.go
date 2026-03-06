package utils

import (
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiter applies a global request limit.
func RateLimiter() fiber.Handler {
    return limiter.New(limiter.Config{
        Max:        100,             // 100 requests
        Expiration: 1 * time.Minute, // per minute
    })
}
