package services

import (
    "encoding/json"
    "log"
    "sync"

    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"

    "github.com/google/uuid"
)

// Max 1000 entries buffered.
const activityBufferSize = 1000

var (
    activityQueue chan *models.ActivityLog
    once          sync.Once
)

// initActivityLogger initializes the async log worker exactly once.
func initActivityLogger() {
    once.Do(func() {
        activityQueue = make(chan *models.ActivityLog, activityBufferSize)

        // Worker goroutine processing logs asynchronously.
        go func() {
            for entry := range activityQueue {
                if err := repositories.InsertActivity(entry); err != nil {
                    log.Printf("Activity log insert failed: %v", err)
                }
            }
        }()
    })
}

// LogActivity queues an activity log entry asynchronously.
func LogActivity(userID uuid.UUID, action string, metadata map[string]interface{}, ip string, userAgent string) {
    initActivityLogger()

    raw, _ := json.Marshal(metadata)

    entry := &models.ActivityLog{
        UserID:    userID,
        Action:    action,
        Metadata:  raw,
        IP:        ip,
        UserAgent: userAgent,
    }

    select {
    case activityQueue <- entry:
    default:
        // If queue is full, drop the event to prevent blocking.
        log.Printf("Activity queue full — dropping event: %s", action)
    }
}

// ShutdownActivityLogger flushes logs on server stop.
func ShutdownActivityLogger() {
    if activityQueue != nil {
        close(activityQueue)
    }
}
