package services

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/google/uuid"
)

// AddFavorite handles the business logic for creating a favorite.
func AddFavorite(userID uuid.UUID, req dto.CreateFavoriteRequest) error {
    fav := &models.UserFavorite{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        Title:           req.Title,
        Preview:         req.Preview,
        DurationSeconds: req.DurationSeconds,
    }
    return repositories.AddFavorite(fav)
}

// RemoveFavorite removes a favorite entry.
func RemoveFavorite(userID uuid.UUID, req dto.RemoveFavoriteRequest) error {
    return repositories.RemoveFavorite(userID, req.ContentID, req.ContentType)
}

// GetFavorites fetches all favorites in newest-first order.
func GetFavorites(userID uuid.UUID) ([]models.UserFavorite, error) {
    return repositories.ListFavorites(userID)
}
