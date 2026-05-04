package dto

// UpdateProgressRequest updates resume info + triggers analytics event.
type UpdateProgressRequest struct {
    ContentID       string `json:"content_id" validate:"required,max=255"`
    ContentType     string `json:"content_type" validate:"required,max=100"`
    PositionSeconds int    `json:"position_seconds" validate:"min=0"`
    DurationSeconds int    `json:"duration_seconds" validate:"min=1"`
    SessionID       string `json:"session_id" validate:"max=36"`
    EventType       string `json:"event_type" validate:"required,oneof=start pause progress complete"`
}

// ClearHistoryRequest deletes all history for a user (optional extension)
type ClearHistoryRequest struct{}
