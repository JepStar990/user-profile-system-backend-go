package controllers

import (
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
)

type AdminVersionController struct{}

func (AdminVersionController) Version(c *fiber.Ctx) error {

    return c.JSON(fiber.Map{
        "version":        os.Getenv("APP_VERSION"),
        "commit":         os.Getenv("GIT_COMMIT"),
        "build_date":     os.Getenv("BUILD_DATE"),
        "timestamp":      time.Now().UTC(),
    })
}
