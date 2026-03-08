package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
    Email        string    `gorm:"uniqueIndex;size:255;not null"`
    Username     string    `gorm:"uniqueIndex;size:255;not null"`
    PasswordHash string    `gorm:"size:255;not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return nil
}
