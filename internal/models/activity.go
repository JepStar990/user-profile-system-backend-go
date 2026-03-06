package models

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// ActivityLog represents one internal-only log entry recording
// actions performed by a user (e.g., login, profile update, downloads, etc.).
type ActivityLog struct {
    ID        uuid.UUID       `gorm:"type:char(36);primaryKey"`
    UserID    uuid.UUID       `gorm:"type:char(36);index;not null"`
    Action    string          `gorm:"size:255;not null"`
    Metadata  json.RawMessage `gorm:"type:json"`
    IP        string          `gorm:"size:100"`
    UserAgent string          `gorm:"size:512"`
    CreatedAt time.Time
}

func (a *ActivityLog) BeforeCreate(tx *gorm.DB) error {
    a.ID = uuid.New()
    return nil
}
