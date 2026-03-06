package services

import (
    "encoding/json"
    "fmt"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// -----------------------
//  SETTINGS SERVICE
// -----------------------

func GetSettings(userID uuid.UUID) (map[string]interface{}, error) {
    s, err := repositories.EnsureSettings(userID, defaultSettings)
    if err != nil {
        return nil, err
    }

    var decoded map[string]interface{}
    err = json.Unmarshal(s.SettingsRaw, &decoded)
    return decoded, err
}

func saveSettings(userID uuid.UUID, settings map[string]interface{}) error {
    raw, err := json.Marshal(settings)
    if err != nil {
        return err
    }
    return repositories.UpdateSettingsJSON(userID, raw)
}

// mergeSection updates only one settings block (audio, voice, etc.)
func mergeSection(userID uuid.UUID, section string, payload interface{}) error {
    settings, err := GetSettings(userID)
    if err != nil {
        return err
    }

    sectionMap, ok := settings[section].(map[string]interface{})
    if !ok {
        return fmt.Errorf("settings section '%s' invalid", section)
    }

    patchBytes, _ := json.Marshal(payload)
    var patch map[string]interface{}
    json.Unmarshal(patchBytes, &patch)

    for key, val := range patch {
        sectionMap[key] = val
    }

    settings[section] = sectionMap
    return saveSettings(userID, settings)
}

// Each update function logs activity:
func UpdateAudioSettings(userID uuid.UUID, req dto.AudioSettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "audio", req)
    if err == nil {
        LogActivity(userID, "update_settings_audio", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func UpdateVoiceSettings(userID uuid.UUID, req dto.VoiceSettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "voice", req)
    if err == nil {
        LogActivity(userID, "update_settings_voice", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func UpdateLiveRadioSettings(userID uuid.UUID, req dto.LiveRadioSettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "live_radio", req)
    if err == nil {
        LogActivity(userID, "update_settings_live_radio", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func UpdateNotificationSettings(userID uuid.UUID, req dto.NotificationSettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "notifications", req)
    if err == nil {
        LogActivity(userID, "update_settings_notifications", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func UpdateAppearanceSettings(userID uuid.UUID, req dto.AppearanceSettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "appearance", req)
    if err == nil {
        LogActivity(userID, "update_settings_appearance", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func UpdatePrivacySettings(userID uuid.UUID, req dto.PrivacySettingsRequest, c *fiber.Ctx) error {
    err := mergeSection(userID, "privacy", req)
    if err == nil {
        LogActivity(userID, "update_settings_privacy", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}
