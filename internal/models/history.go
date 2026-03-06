package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// UserHistory tracks the last progress for each content item.
// One row per (user + content + type).
type UserHistory struct {
    ID                 uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID             uuid.UUID `gorm:"type:char(36);index;not null"`
    ContentID          string    `gorm:"size:255;not null"`
    ContentType        string    `gorm:"size:100;not null"`
    LastPositionSeconds int       `gorm:"default:0"`
    DurationSeconds    int       `gorm:"default:0"`
    Completed          bool      `gorm:"default:false"`
    UpdatedAt          time.Time
}

func (h *UserHistory) BeforeCreate(tx *gorm.DB) error {
    h.ID = uuid.New()
    return nil
}
