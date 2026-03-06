package utils

import (
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
)

// RequestLogger logs structured request information.
func RequestLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()
        err := c.Next()
        latency := time.Since(start)

        reqID := c.Locals("request_id")
        userID := c.Locals("user_id")

        log.Printf(
            "[REQ] id=%v user=%v %s %s %d %v IP=%s UA=\"%s\"",
            reqID,
            userID,
            c.Method(),
            c.Path(),
            c.Response().StatusCode(),
            latency,
            c.IP(),
            string(c.Context().UserAgent()),
        )

        return err
    }
}
