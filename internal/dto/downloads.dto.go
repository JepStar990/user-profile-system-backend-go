package dto

// CreateDownloadRequest is used to register a download entry.
type CreateDownloadRequest struct {
    ContentID       string `json:"content_id" validate:"required"`
    ContentType     string `json:"content_type" validate:"required"`
    DownloadQuality string `json:"download_quality" validate:"required,oneof=low medium high"`
}

// RemoveDownloadRequest deletes one download entry.
type RemoveDownloadRequest struct {
    ContentID   string `json:"content_id" validate:"required"`
    ContentType string `json:"content_type" validate:"required"`
}

// PresignedURLResponse is returned when frontend requests a presigned URL.
type PresignedURLResponse struct {
    URL string `json:"url"`
}
