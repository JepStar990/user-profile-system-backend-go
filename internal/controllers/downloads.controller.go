package controllers

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type DownloadsController struct{}

// POST /downloads
func (DownloadsController) RegisterDownload(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.CreateDownloadRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(body); err != nil {
        return err
    }

    if err := services.RegisterDownload(userID, body); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Download registered"})
}

// GET /downloads
func (DownloadsController) ListDownloads(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    items, err := services.GetDownloads(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"downloads": items})
}

// DELETE /downloads
func (DownloadsController) RemoveDownload(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.RemoveDownloadRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(body); err != nil {
        return err
    }

    if err := services.RemoveDownload(userID, body.ContentID, body.ContentType); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Download removed"})
}

// GET /downloads/url?content_id=...&content_type=...
func (DownloadsController) PresignedURL(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    contentID := c.Query("content_id")
    contentType := c.Query("content_type")

    url, err := services.GetPresignedURL(userID, contentID, contentType)
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, err.Error())
    }

    return c.JSON(dto.PresignedURLResponse{URL: url})
}
