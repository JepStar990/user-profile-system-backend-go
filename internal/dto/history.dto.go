package dto

// UpdateProgressRequest updates resume info + triggers analytics event.
type UpdateProgressRequest struct {
    ContentID       string `json:"content_id" validate:"required"`
    ContentType     string `json:"content_type" validate:"required"`
    PositionSeconds int    `json:"position_seconds" validate:"min=0"`
    DurationSeconds int    `json:"duration_seconds" validate:"min=1"`
    SessionID       string `json:"session_id"` // optional; backend generates if missing
    EventType       string `json:"event_type" validate:"required,oneof=start pause progress complete"`
}

// ClearHistoryRequest deletes all history for a user (optional extension)
type ClearHistoryRequest struct{}
