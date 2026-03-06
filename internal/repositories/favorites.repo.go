package repositories

import (
    "errors"

    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// AddFavorite inserts a new favorite. Duplicate = no error because of our upsert logic.
func AddFavorite(fav *models.UserFavorite) error {
    err := db.DB.Create(fav).Error
    if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
        // Already exists — idempotent behavior
        return nil
    }
    return err
}

// RemoveFavorite deletes a record based on user/content IDs.
func RemoveFavorite(userID uuid.UUID, contentID, contentType string) error {
    return db.DB.
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        Delete(&models.UserFavorite{}).
        Error
}

// ListFavorites returns favorites newest-first.
func ListFavorites(userID uuid.UUID) ([]models.UserFavorite, error) {
    var items []models.UserFavorite
    err := db.DB.
        Where("user_id = ?", userID).
        Order("created_at DESC").
        Find(&items).
        Error
    return items, err
}

// Exists checks if a favorite exists
func Exists(userID uuid.UUID, contentID, contentType string) bool {
    var count int64
    db.DB.Model(&models.UserFavorite{}).
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        Count(&count)
    return count > 0
}
