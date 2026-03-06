package controllers

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type FavoritesController struct{}

// AddFavorite  → POST /favorites
func (FavoritesController) AddFavorite(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.CreateFavoriteRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(body); err != nil {
        return err
    }

    err := services.AddFavorite(userID, body)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Added to favorites"})
}

// RemoveFavorite  → DELETE /favorites
func (FavoritesController) RemoveFavorite(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.RemoveFavoriteRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(body); err != nil {
        return err
    }

    err := services.RemoveFavorite(userID, body)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Removed from favorites"})
}

// ListFavorites  → GET /favorites
func (FavoritesController) ListFavorites(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    items, err := services.GetFavorites(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{
        "favorites": items,
    })
}
