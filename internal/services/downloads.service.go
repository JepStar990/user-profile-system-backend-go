package services

import (
    "errors"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/storage"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// -----------------------
//  DOWNLOADS SERVICE
// -----------------------

func RegisterDownload(userID uuid.UUID, req dto.CreateDownloadRequest, c *fiber.Ctx) error {
    d := &models.UserDownload{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        DownloadQuality: req.DownloadQuality,
        StorageURL:      "media/" + req.ContentID + ".mp3",
    }

    err := repositories.CreateOrUpdate(d)
    if err == nil {
        LogActivity(userID, "register_download", map[string]any{
            "content_id":   req.ContentID,
            "content_type": req.ContentType,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}

func GetDownloads(userID uuid.UUID) ([]models.UserDownload, error) {
    return repositories.ListDownloads(userID)
}

func RemoveDownload(userID uuid.UUID, contentID, contentType string, c *fiber.Ctx) error {
    err := repositories.DeleteDownload(userID, contentID, contentType)
    if err == nil {
        LogActivity(userID, "remove_download", map[string]any{
            "content_id":   contentID,
            "content_type": contentType,
        }, c.IP(), string(c.Context().UserAgent()))
    }
    return err
}

func GetPresignedURL(userID uuid.UUID, contentID, contentType string, c *fiber.Ctx) (string, error) {
    d, err := repositories.FindDownload(userID, contentID, contentType)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return "", errors.New("download entry not found")
    }

    presigner := storage.NewS3Presigner()
    url, err := presigner.GenerateDownloadURL(d.StorageURL)

    if err == nil {
        LogActivity(userID, "download_presigned_url", map[string]any{
            "content_id":   contentID,
            "content_type": contentType,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return url, err
}
