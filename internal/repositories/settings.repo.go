package repositories

import (
    "encoding/json"
    "errors"
    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// GetSettings returns settings row or gorm.ErrRecordNotFound.
func GetSettings(userID uuid.UUID) (*models.UserSettings, error) {
    var s models.UserSettings
    err := db.DB.First(&s, "user_id = ?", userID).Error
    return &s, err
}

// CreateDefaultSettings inserts a fresh settings JSON row.
func CreateDefaultSettings(userID uuid.UUID, defaults map[string]interface{}) error {
    raw, err := json.Marshal(defaults)
    if err != nil {
        return err
    }

    settings := models.UserSettings{
        UserID:      userID,
        SettingsRaw: raw,
    }

    return db.DB.Create(&settings).Error
}

// UpdateSettingsJSON replaces the entire settings JSON.
func UpdateSettingsJSON(userID uuid.UUID, newJSON []byte) error {
    return db.DB.Model(&models.UserSettings{}).
        Where("user_id = ?", userID).
        Update("settings_raw", newJSON).Error
}

// EnsureSettings returns existing settings or auto-creates defaults.
func EnsureSettings(userID uuid.UUID, defaults map[string]interface{}) (*models.UserSettings, error) {
    s, err := GetSettings(userID)

    if errors.Is(err, gorm.ErrRecordNotFound) {
        err = CreateDefaultSettings(userID, defaults)
        if err != nil {
            return nil, err
        }
        return GetSettings(userID)
    }

    return s, err
}
