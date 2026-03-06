package controllers

import (
    "runtime"
    "time"
    "user-profile-system-backend-go/internal/db"

    "github.com/gofiber/fiber/v2"
)

type AdminHealthController struct {
    start time.Time
}

func NewAdminHealthController() *AdminHealthController {
    return &AdminHealthController{
        start: time.Now(),
    }
}

func (ctl *AdminHealthController) Health(c *fiber.Ctx) error {

    // DB health check
    sqlDB, err := db.DB.DB()
    dbStatus := "ok"
    if err != nil {
        dbStatus = "error: cannot access DB"
    } else if pingErr := sqlDB.Ping(); pingErr != nil {
        dbStatus = "error: ping failed"
    }

    return c.JSON(fiber.Map{
        "status":          "ok",
        "uptime_seconds":  int(time.Since(ctl.start).Seconds()),
        "goroutines":      runtime.NumGoroutine(),
        "go_version":      runtime.Version(),
        "database_status": dbStatus,
        "timestamp":       time.Now().UTC(),
    })
}
