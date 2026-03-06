package models

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
)

// UserSettings represents all settings stored in a single JSON column.
type UserSettings struct {
    UserID      uuid.UUID       `gorm:"type:char(36);primaryKey"`
    SettingsRaw json.RawMessage `gorm:"type:json;not null"`
    UpdatedAt   time.Time
}
