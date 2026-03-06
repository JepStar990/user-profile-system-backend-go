package services

import (
    "time"

    "user-profile-system-backend-go/internal/repositories"

    "github.com/google/uuid"
)

// StatsResponse returned to frontend
type StatsResponse struct {
    TotalListeningSeconds int64                  `json:"total_listening_seconds"`
    CompletedItems        int                    `json:"completed_items"`
    DailyMinutes          map[string]int64       `json:"daily_minutes"`
    TopContentTypes       map[string]int         `json:"top_content_types"`
}

// ComputeStats builds analytics metrics for the user.
func ComputeStats(userID uuid.UUID) (*StatsResponse, error) {

    // We fetch events from last 30 days by default
    since := time.Now().Add(-30 * 24 * time.Hour)

    events, err := repositories.GetEventsSince(userID, since)
    if err != nil {
        return nil, err
    }

    totalSeconds := int64(0)
    completed := 0
    topTypes := map[string]int{}
    daily := map[string]int64{}

    // Compute per-event deltas
    for i := 1; i < len(events); i++ {
        prev := events[i-1]
        curr := events[i]

        // Only count progress deltas inside same session
        if prev.SessionID == curr.SessionID &&
            curr.PositionSeconds > prev.PositionSeconds &&
            curr.EventType == "progress" {

            delta := int64(curr.PositionSeconds - prev.PositionSeconds)
            totalSeconds += delta

            dayKey := curr.CreatedAt.Format("2006-01-02")
            daily[dayKey] += delta
        }

        if curr.EventType == "complete" {
            completed++
        }

        // Count content types
        topTypes[curr.ContentType]++
    }

    return &StatsResponse{
        TotalListeningSeconds: totalSeconds,
        CompletedItems:        completed,
        DailyMinutes:          daily,
        TopContentTypes:       topTypes,
    }, nil
}
