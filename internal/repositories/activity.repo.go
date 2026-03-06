package repositories

import (
    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"
)

// InsertActivity writes a log entry into DB.
func InsertActivity(log *models.ActivityLog) error {
    return db.DB.Create(log).Error
}
