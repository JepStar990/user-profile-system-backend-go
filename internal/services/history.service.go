package services

import (
    "errors"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// -----------------------
//  HISTORY SERVICE
// -----------------------

func UpdateProgress(userID uuid.UUID, req dto.UpdateProgressRequest, c *fiber.Ctx) error {

    // Session ID
    sessionID := uuid.New()
    if req.SessionID != "" {
        parsed, err := uuid.Parse(req.SessionID)
        if err == nil {
            sessionID = parsed
        }
    }

    // Resume history
    h, err := repositories.GetSingleHistory(userID, req.ContentID, req.ContentType)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        h = &models.UserHistory{
            UserID:             userID,
            ContentID:          req.ContentID,
            ContentType:        req.ContentType,
            DurationSeconds:    req.DurationSeconds,
            LastPositionSeconds: req.PositionSeconds,
        }
    } else {
        h.LastPositionSeconds = req.PositionSeconds
        h.DurationSeconds = req.DurationSeconds
        if req.PositionSeconds >= int(float64(req.DurationSeconds)*0.9) {
            h.Completed = true
        }
    }

    if err := repositories.SaveHistory(h); err != nil {
        return err
    }

    // Insert analytics event
    e := &models.ListeningEvent{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        SessionID:       sessionID,
        EventType:       req.EventType,
        PositionSeconds: req.PositionSeconds,
        DurationSeconds: req.DurationSeconds,
    }
    err = repositories.InsertEvent(e)

    if err == nil {
        LogActivity(userID, "history_progress", map[string]any{
            "content_id":   req.ContentID,
            "content_type": req.ContentType,
            "event":        req.EventType,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}

func GetHistory(userID uuid.UUID) ([]models.UserHistory, error) {
    return repositories.GetHistory(userID)
}

func ClearUserHistory(userID uuid.UUID, c *fiber.Ctx) error {
    err := repositories.ClearHistory(userID)

    if err == nil {
        LogActivity(userID, "clear_history", nil, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}
