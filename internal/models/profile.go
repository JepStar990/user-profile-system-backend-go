package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type UserProfile struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID    uuid.UUID `gorm:"type:char(36);uniqueIndex;not null"`
    FullName  string    `gorm:"size:255"`
    AvatarURL string    `gorm:"size:512"`
    Bio       string    `gorm:"type:text"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (p *UserProfile) BeforeCreate(tx *gorm.DB) error {
    p.ID = uuid.New()
    return nil
}
