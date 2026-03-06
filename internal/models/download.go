package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// UserDownload represents a downloaded item.
// One download entry per (user + content + type).
type UserDownload struct {
    ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
    UserID          uuid.UUID `gorm:"type:char(36);index;not null"`
    ContentID       string    `gorm:"size:255;not null"`
    ContentType     string    `gorm:"size:100;not null"`
    DownloadQuality string    `gorm:"size:50;not null"`
    FileSizeBytes   int64     `gorm:"default:0"`
    StorageURL      string    `gorm:"size:512;not null"`
    Status          string    `gorm:"size:50;default:'ready'"`
    CreatedAt       time.Time
}

func (d *UserDownload) BeforeCreate(tx *gorm.DB) error {
    d.ID = uuid.New()
    return nil
}
