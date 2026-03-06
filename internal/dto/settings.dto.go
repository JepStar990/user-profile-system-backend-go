package dto

//
// AUDIO SETTINGS
//
type AudioSettingsRequest struct {
    Autoplay       bool    `json:"autoplay"`
    BackgroundPlay bool    `json:"background_play"`
    PlaybackSpeed  float64 `json:"playback_speed" validate:"min=0.5,max=2.0"`
    DownloadQuality string `json:"download_quality" validate:"oneof=low medium high"`
}

//
// VOICE SETTINGS
//
type VoiceSettingsRequest struct {
    Enabled  bool   `json:"enabled"`
    WakeWord string `json:"wake_word"` // null or string
}

//
// LIVE RADIO SETTINGS
//
type LiveRadioSettingsRequest struct {
    LiveDefault     bool     `json:"live_default"`
    AutoSkipSilence bool     `json:"auto_skip_silence"`
    Schedule        []string `json:"schedule"`
}

//
// NOTIFICATIONS SETTINGS
//
type NotificationSettingsRequest struct {
    BreakingNews bool `json:"breaking_news"`
    Categories   struct {
        Breaking   bool `json:"breaking"`
        Politics   bool `json:"politics"`
        Technology bool `json:"technology"`
        Business   bool `json:"business"`
    } `json:"categories"`
}

//
// APPEARANCE SETTINGS
//
type AppearanceSettingsRequest struct {
    DarkMode          bool   `json:"dark_mode"`
    ThemeColor        string `json:"theme_color"`
    BackgroundOpacity int    `json:"background_opacity" validate:"min=0,max=100"`
    BlurEffects       bool   `json:"blur_effects"`
}

//
// PRIVACY SETTINGS
//
type PrivacySettingsRequest struct {
    AllowDataCollection bool `json:"allow_data_collection"`
}

//
// FULL RESPONSE DTO
//
type SettingsResponse struct {
    Settings map[string]interface{} `json:"settings"`
}
