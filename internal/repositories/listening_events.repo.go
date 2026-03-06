package repositories

import (
    "time"

    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
)

// InsertEvent inserts a single analytics event.
func InsertEvent(e *models.ListeningEvent) error {
    return db.DB.Create(e).Error
}

// GetEvents returns all events for stats.
func GetEvents(userID uuid.UUID) ([]models.ListeningEvent, error) {
    var items []models.ListeningEvent
    err := db.DB.
        Where("user_id = ?", userID).
        Order("created_at DESC").
        Find(&items).Error
    return items, err
}

// GetEventsSince returns events for daily stats graph.
func GetEventsSince(userID uuid.UUID, since time.Time) ([]models.ListeningEvent, error) {
    var items []models.ListeningEvent
    err := db.DB.
        Where("user_id = ? AND created_at >= ?", userID, since).
        Order("created_at ASC").
        Find(&items).Error
    return items, err
}
