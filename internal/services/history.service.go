package services

import (
    "errors"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// UpdateProgress updates resume history AND logs event in analytics.
func UpdateProgress(userID uuid.UUID, req dto.UpdateProgressRequest) error {

    // -------- 1. Parse or generate session ID --------
    sessionID := uuid.New()
    if req.SessionID != "" {
        parsed, err := uuid.Parse(req.SessionID)
        if err == nil {
            sessionID = parsed
        }
    }

    // -------- 2. Update resume table (user_history) --------
    h, err := repositories.GetSingleHistory(userID, req.ContentID, req.ContentType)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        h = &models.UserHistory{
            UserID:           userID,
            ContentID:        req.ContentID,
            ContentType:      req.ContentType,
            DurationSeconds:  req.DurationSeconds,
            LastPositionSeconds: req.PositionSeconds,
            Completed:        false,
        }
    } else {
        // Update resume fields
        h.LastPositionSeconds = req.PositionSeconds
        h.DurationSeconds = req.DurationSeconds

        // Mark completed
        if req.PositionSeconds >= int(float64(req.DurationSeconds)*0.9) {
            h.Completed = true
        }
    }

    if err := repositories.SaveHistory(h); err != nil {
        return err
    }

    // -------- 3. Insert analytics event (user_listening_events) --------
    e := &models.ListeningEvent{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        SessionID:       sessionID,
        EventType:       req.EventType,
        PositionSeconds: req.PositionSeconds,
        DurationSeconds: req.DurationSeconds,
    }

    return repositories.InsertEvent(e)
}

func GetHistory(userID uuid.UUID) ([]models.UserHistory, error) {
    return repositories.GetHistory(userID)
}

func ClearUserHistory(userID uuid.UUID) error {
    return repositories.ClearHistory(userID)
}
