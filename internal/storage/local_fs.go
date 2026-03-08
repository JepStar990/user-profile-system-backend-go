package storage

import (
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// LocalFSUploader stores uploaded files on the local filesystem.
// Use for local development only (not recommended for production).
type LocalFSUploader struct {
    BaseDir  string // e.g. "./workdir/uploads"
    BaseURL  string // e.g. "http://localhost:8080"
    SubDir   string // e.g. "avatars"
    MaxBytes int64  // safety limit
}

func NewLocalFSUploader() *LocalFSUploader {
    baseDir := os.Getenv("LOCAL_UPLOAD_DIR")
    if baseDir == "" {
        baseDir = "./workdir/uploads"
    }
    baseURL := os.Getenv("LOCAL_BASE_URL")
    if baseURL == "" {
        baseURL = "http://localhost:8080"
    }
    subDir := os.Getenv("LOCAL_UPLOAD_SUBDIR")
    if subDir == "" {
        subDir = "avatars"
    }

    return &LocalFSUploader{
        BaseDir:  baseDir,
        BaseURL:  strings.TrimRight(baseURL, "/"),
        SubDir:   subDir,
        MaxBytes: 5 << 20, // 5MB default max
    }
}

// UploadAvatar saves the file to disk and returns a URL that can be served by your app.
// NOTE: You must expose the BaseDir over HTTP if you want the URL to work externally.
func (u *LocalFSUploader) UploadAvatar(file *multipart.FileHeader, userID string) (string, error) {
    if file == nil {
        return "", fmt.Errorf("file is required")
    }
    if file.Size > u.MaxBytes {
        return "", fmt.Errorf("file too large: %d bytes (max %d)", file.Size, u.MaxBytes)
    }

    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // Create target directory: BaseDir/SubDir
    targetDir := filepath.Join(u.BaseDir, u.SubDir)
    if err := os.MkdirAll(targetDir, 0o755); err != nil {
        return "", err
    }

    // Build safe filename
    ext := filepath.Ext(file.Filename)
    if ext == "" {
        ext = ".bin"
    }
    filename := fmt.Sprintf("%s_%d%s", userID, time.Now().UnixNano(), ext)
    targetPath := filepath.Join(targetDir, filename)

    dst, err := os.Create(targetPath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return "", err
    }

    // URL assumes you serve /uploads from BaseDir (example)
    // If you don’t serve this path, you can still return the path itself.
    url := fmt.Sprintf("%s/uploads/%s/%s", u.BaseURL, u.SubDir, filename)
    return url, nil
}
