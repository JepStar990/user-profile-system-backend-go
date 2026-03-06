package services

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// -----------------------
//  FAVORITES SERVICE
// -----------------------

func AddFavorite(userID uuid.UUID, req dto.CreateFavoriteRequest, c *fiber.Ctx) error {
    fav := &models.UserFavorite{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        Title:           req.Title,
        Preview:         req.Preview,
        DurationSeconds: req.DurationSeconds,
    }

    err := repositories.AddFavorite(fav)
    if err == nil {
        LogActivity(userID, "add_favorite", map[string]any{
            "content_id":   req.ContentID,
            "content_type": req.ContentType,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}

func RemoveFavorite(userID uuid.UUID, req dto.RemoveFavoriteRequest, c *fiber.Ctx) error {
    err := repositories.RemoveFavorite(userID, req.ContentID, req.ContentType)
    if err == nil {
        LogActivity(userID, "remove_favorite", map[string]any{
            "content_id":   req.ContentID,
            "content_type": req.ContentType,
        }, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func GetFavorites(userID uuid.UUID) ([]models.UserFavorite, error) {
    return repositories.ListFavorites(userID)
}
