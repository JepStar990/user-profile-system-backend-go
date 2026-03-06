package controllers

import (
    "runtime"
    "time"

    "github.com/gofiber/fiber/v2"
)

type AdminMetricsController struct{}

func (AdminMetricsController) Metrics(c *fiber.Ctx) error {
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)

    return c.JSON(fiber.Map{
        "timestamp": time.Now().UTC(),
        "memory": fiber.Map{
            "alloc":       mem.Alloc,
            "total_alloc": mem.TotalAlloc,
            "sys":         mem.Sys,
            "heap_alloc":  mem.HeapAlloc,
            "heap_sys":    mem.HeapSys,
        },
        "gc": fiber.Map{
            "num_gc": mem.NumGC,
        },
        "goroutines": runtime.NumGoroutine(),
    })
}
