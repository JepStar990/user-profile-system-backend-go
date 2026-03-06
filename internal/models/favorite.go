package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// UserFavorite represents a single saved favorite item for a user.
// Uniqueness: user_id + content_id + content_type.
type UserFavorite struct {
    ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID          uuid.UUID `gorm:"type:char(36);index;not null"`
    ContentID       string    `gorm:"size:255;not null"`
    ContentType     string    `gorm:"size:100;not null"`
    Title           string    `gorm:"size:255"`
    Preview         string    `gorm:"type:text"`
    DurationSeconds int       `gorm:"default:0"`
    CreatedAt       time.Time
}

func (f *UserFavorite) BeforeCreate(tx *gorm.DB) error {
    f.ID = uuid.New()
    return nil
}
