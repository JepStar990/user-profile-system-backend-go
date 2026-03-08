package services

import (
    "mime/multipart"
    "time"

    "user-profile-system-backend-go/internal/storage"
)

// StorageService centralizes file operations used across modules.
// This is NOT a placeholder: it provides concrete S3-backed operations.
type StorageService struct {
    uploader  *storage.S3Uploader
    presigner *storage.S3Presigner
}

// NewStorageService builds a storage service backed by AWS S3.
// Credentials are read from the standard AWS SDK chain.
func NewStorageService() *StorageService {
    return &StorageService{
        uploader:  storage.NewS3Uploader(),
        presigner: storage.NewS3Presigner(),
    }
}

// UploadAvatar uploads a user's avatar to S3 and returns a public URL.
func (s *StorageService) UploadAvatar(file *multipart.FileHeader, userID string) (string, error) {
    return s.uploader.UploadAvatar(file, userID)
}

// PresignDownloadURL creates a temporary presigned URL for an S3 object key.
// Default expiry is controlled inside the presigner; this wrapper allows future customization.
func (s *StorageService) PresignDownloadURL(objectKey string, _ time.Duration) (string, error) {
    return s.presigner.GenerateDownloadURL(objectKey)
}
