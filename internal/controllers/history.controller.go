package controllers

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type HistoryController struct{}

// GET /history
func (HistoryController) GetHistory(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    history, err := services.GetHistory(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }
    return c.JSON(fiber.Map{"history": history})
}

// POST /history/progress
func (HistoryController) UpdateProgress(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.UpdateProgressRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(body); err != nil {
        return err
    }

    if err := services.UpdateProgress(userID, body); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Progress updated"})
}

// DELETE /history
func (HistoryController) ClearHistory(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    if err := services.ClearUserHistory(userID); err != nil {
        return fiber.ErrInternalServerError
    }
    return c.JSON(fiber.Map{"message": "History cleared"})
}

// GET /history/stats
func (HistoryController) Stats(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    stats, err := services.ComputeStats(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"stats": stats})
}
