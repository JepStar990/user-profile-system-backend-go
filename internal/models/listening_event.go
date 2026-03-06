package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// ListeningEvent captures analytics events like start, pause, progress, complete.
type ListeningEvent struct {
    ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID          uuid.UUID `gorm:"type:char(36);index;not null"`
    ContentID       string    `gorm:"size:255;not null"`
    ContentType     string    `gorm:"size:100;not null"`
    SessionID       uuid.UUID `gorm:"type:char(36);not null"`
    EventType       string    `gorm:"size:50;not null"` // start | pause | progress | complete
    PositionSeconds int       `gorm:"default:0"`
    DurationSeconds int       `gorm:"default:0"`
    CreatedAt       time.Time
}

func (e *ListeningEvent) BeforeCreate(tx *gorm.DB) error {
    e.ID = uuid.New()
    return nil
}
