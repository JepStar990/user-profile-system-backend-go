package services

import (
    "errors"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/storage"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// Register a download (metadata only)
func RegisterDownload(userID uuid.UUID, req dto.CreateDownloadRequest) error {
    d := &models.UserDownload{
        UserID:          userID,
        ContentID:       req.ContentID,
        ContentType:     req.ContentType,
        DownloadQuality: req.DownloadQuality,

        // In a real system, you would fetch metadata like file size from S3.
        FileSizeBytes: 0,
        Status:        "ready",

        // Permanent storage location — this is where the file lives in S3
        StorageURL: "media/" + req.ContentID + ".mp3", // Example — adjust later
    }

    return repositories.CreateOrUpdate(d)
}

// Get download entries
func GetDownloads(userID uuid.UUID) ([]models.UserDownload, error) {
    return repositories.ListDownloads(userID)
}

func RemoveDownload(userID uuid.UUID, contentID, contentType string) error {
    return repositories.DeleteDownload(userID, contentID, contentType)
}

func GetPresignedURL(userID uuid.UUID, contentID, contentType string) (string, error) {
    d, err := repositories.FindDownload(userID, contentID, contentType)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return "", errors.New("download entry not found")
    }

    presigner := storage.NewS3Presigner()

    // d.StorageURL is the key in S3
    return presigner.GenerateDownloadURL(d.StorageURL)
}
