package dto

//
// AUDIO SETTINGS
//
type AudioSettingsRequest struct {
    Autoplay        bool    `json:"autoplay"`
    BackgroundPlay  bool    `json:"background_play"`
    PlaybackSpeed   float64 `json:"playback_speed" validate:"omitempty,min=0.5,max=2.0"`
    DownloadQuality string  `json:"download_quality" validate:"omitempty,oneof=low medium high "`
    Volume          float64 `json:"volume" validate:"omitempty,min=0,max=1"`
}

//
// VOICE SETTINGS
//
type VoiceSettingsRequest struct {
    Enabled  bool    `json:"enabled"`
    WakeWord string  `json:"wake_word" validate:"max=50"`
    Speed    float64 `json:"speed" validate:"omitempty,min=0.5,max=2.0"`
    Language string  `json:"language" validate:"omitempty,max=10"`
}

//
// LIVE RADIO SETTINGS
//
type LiveRadioSettingsRequest struct {
    LiveDefault     bool     `json:"live_default"`
    AutoSkipSilence bool     `json:"auto_skip_silence"`
    AutoPlay        bool     `json:"auto_play"`
    Schedule        []string `json:"schedule"`
}

//
// NOTIFICATIONS SETTINGS
//
type NotificationSettingsRequest struct {
    PushEnabled  bool `json:"push_enabled"`
    EmailEnabled bool `json:"email_enabled"`
    BreakingNews bool `json:"breaking_news"`
}

//
// APPEARANCE SETTINGS
//
type AppearanceSettingsRequest struct {
    DarkMode          bool   `json:"dark_mode"`
    Theme             string `json:"theme" validate:"omitempty,max=20"`
    ThemeColor        string `json:"theme_color" validate:"omitempty,max=20"`
    BackgroundOpacity int    `json:"background_opacity" validate:"omitempty,min=0,max=100"`
    BlurEffects       bool   `json:"blur_effects"`
    FontSize          string `json:"font_size" validate:"omitempty,max=10"`
}

//
// PRIVACY SETTINGS
//
type PrivacySettingsRequest struct {
    ShowActivity        bool `json:"show_activity"`
    ShowFavorites       bool `json:"show_favorites"`
    AllowDataCollection bool `json:"allow_data_collection"`
}

//
// FULL RESPONSE DTO
//
type SettingsResponse struct {
    Settings map[string]interface{} `json:"settings"`
}
