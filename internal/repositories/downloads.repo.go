package repositories

import (
    "errors"

    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// CreateOrUpdate stores metadata if new, otherwise returns success.
func CreateOrUpdate(d *models.UserDownload) error {
    err := db.DB.Create(d).Error
    if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
        // Already exists → treat as success (idempotent)
        return nil
    }
    return err
}

func DeleteDownload(userID uuid.UUID, contentID, contentType string) error {
    return db.DB.
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        Delete(&models.UserDownload{}).
        Error
}

func ListDownloads(userID uuid.UUID) ([]models.UserDownload, error) {
    var items []models.UserDownload
    err := db.DB.
        Where("user_id = ?", userID).
        Order("created_at DESC").
        Find(&items).
        Error
    return items, err
}

func FindDownload(userID uuid.UUID, contentID, contentType string) (*models.UserDownload, error) {
    var d models.UserDownload
    err := db.DB.
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        First(&d).Error
    return &d, err
}

func Exists(userID uuid.UUID, contentID, contentType string) bool {
    var count int64
    db.DB.Model(&models.UserDownload{}).
        Where("user_id = ? AND content_id = ? AND content_type = ?", userID, contentID, contentType).
        Count(&count)
    return count > 0
}
