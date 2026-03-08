package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type RefreshToken struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID    uuid.UUID `gorm:"type:char(36);index;not null"`
    TokenHash string    `gorm:"size:255;not null"`
    ExpiresAt time.Time `gorm:"index"`
    Revoked   bool      `gorm:"default:false"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) error {
    rt.ID = uuid.New()
    return nil
}
