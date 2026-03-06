package services

import (
    "encoding/json"
    "fmt"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/google/uuid"
)

// Default settings (all disabled — privacy first).
var defaultSettings = map[string]interface{}{
    "audio": map[string]interface{}{
        "autoplay":         false,
        "background_play":  false,
        "playback_speed":   1.0,
        "download_quality": "low",
    },
    "voice": map[string]interface{}{
        "enabled":   false,
        "wake_word": nil,
    },
    "live_radio": map[string]interface{}{
        "live_default":     false,
        "auto_skip_silence": false,
        "schedule":         []string{},
    },
    "notifications": map[string]interface{}{
        "breaking_news": false,
        "categories": map[string]interface{}{
            "breaking":   false,
            "politics":   false,
            "technology": false,
            "business":   false,
        },
    },
    "appearance": map[string]interface{}{
        "dark_mode":          false,
        "theme_color":        nil,
        "background_opacity": 100,
        "blur_effects":       false,
    },
    "privacy": map[string]interface{}{
        "allow_data_collection": false,
    },
}

// GetSettings returns full JSON settings for the user.
func GetSettings(userID uuid.UUID) (map[string]interface{}, error) {
    s, err := repositories.EnsureSettings(userID, defaultSettings)
    if err != nil {
        return nil, err
    }

    var decoded map[string]interface{}
    err = json.Unmarshal(s.SettingsRaw, &decoded)
    return decoded, err
}

// saveSettings writes JSON back to DB.
func saveSettings(userID uuid.UUID, settings map[string]interface{}) error {
    raw, err := json.Marshal(settings)
    if err != nil {
        return err
    }
    return repositories.UpdateSettingsJSON(userID, raw)
}

// mergeSection updates only a specific part of settings JSON.
func mergeSection(userID uuid.UUID, section string, payload interface{}) error {
    settings, err := GetSettings(userID)
    if err != nil {
        return err
    }

    sectionMap, ok := settings[section].(map[string]interface{})
    if !ok {
        return fmt.Errorf("settings section '%s' invalid or missing", section)
    }

    // Convert struct -> map[string]interface{}
    patchBytes, _ := json.Marshal(payload)

    var patch map[string]interface{}
    json.Unmarshal(patchBytes, &patch)

    // Merge keys
    for key, val := range patch {
        sectionMap[key] = val
    }

    settings[section] = sectionMap

    return saveSettings(userID, settings)
}

//
// Public update methods below
//

func UpdateAudioSettings(userID uuid.UUID, req dto.AudioSettingsRequest) error {
    return mergeSection(userID, "audio", req)
}

func UpdateVoiceSettings(userID uuid.UUID, req dto.VoiceSettingsRequest) error {
    return mergeSection(userID, "voice", req)
}

func UpdateLiveRadioSettings(userID uuid.UUID, req dto.LiveRadioSettingsRequest) error {
    return mergeSection(userID, "live_radio", req)
}

func UpdateNotificationSettings(userID uuid.UUID, req dto.NotificationSettingsRequest) error {
    return mergeSection(userID, "notifications", req)
}

func UpdateAppearanceSettings(userID uuid.UUID, req dto.AppearanceSettingsRequest) error {
    return mergeSection(userID, "appearance", req)
}

func UpdatePrivacySettings(userID uuid.UUID, req dto.PrivacySettingsRequest) error {
    return mergeSection(userID, "privacy", req)
}
