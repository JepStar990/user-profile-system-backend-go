package dto

// CreateFavoriteRequest is the payload used when adding a favorite.
type CreateFavoriteRequest struct {
    ContentID       string `json:"content_id" validate:"required"`
    ContentType     string `json:"content_type" validate:"required"`
    Title           string `json:"title"`
    Preview         string `json:"preview"`
    DurationSeconds int    `json:"duration_seconds" validate:"min=0"`
}

// RemoveFavoriteRequest is used when removing a favorite.
type RemoveFavoriteRequest struct {
    ContentID   string `json:"content_id" validate:"required"`
    ContentType string `json:"content_type" validate:"required"`
}
