package repositories

import (
    "errors"

    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// GetHistory returns all history entries for a user.
func GetHistory(userID uuid.UUID) ([]models.UserHistory, error) {
    var items []models.UserHistory
    err := db.DB.
        Where("user_id = ?", userID).
        Order("updated_at DESC").
        Find(&items).Error
    return items, err
}

// GetSingleHistory returns one record or ErrRecordNotFound.
func GetSingleHistory(userID uuid.UUID, contentID, contentType string) (*models.UserHistory, error) {
    var h models.UserHistory
    err := db.DB.
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        First(&h).Error
    return &h, err
}

func SaveHistory(h *models.UserHistory) error {
    return db.DB.Save(h).Error
}

func ClearHistory(userID uuid.UUID) error {
    return db.DB.
        Where("user_id = ?", userID).
        Delete(&models.UserHistory{}).
        Error
}
