package dto

// CreateFavoriteRequest is the payload used when adding a favorite.
type CreateFavoriteRequest struct {
    ContentID       string `json:"content_id" validate:"required,max=255"`
    ContentType     string `json:"content_type" validate:"required,max=100"`
    Title           string `json:"title" validate:"max=500"`
    Preview         string `json:"preview" validate:"max=2000"`
    DurationSeconds int    `json:"duration_seconds" validate:"min=0"`
}

// RemoveFavoriteRequest is used when removing a favorite.
type RemoveFavoriteRequest struct {
    ContentID   string `json:"content_id" validate:"required,max=255"`
    ContentType string `json:"content_type" validate:"required,max=100"`
}
