package models

type ContentItem struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	DurationSec int      `json:"duration_sec"`
}
