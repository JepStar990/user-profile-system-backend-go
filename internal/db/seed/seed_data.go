package seed

import (
    "os"

    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/security"

    "gorm.io/gorm"
)

// Seed optionally inserts development data.
// Controlled by SEED_ENABLED=true and SEED_EMAIL/SEED_PASSWORD env vars.
// This is safe to keep in repo; it does not run unless you call it explicitly.
func Seed(db *gorm.DB) error {
    if os.Getenv("SEED_ENABLED") != "true" {
        return nil
    }

    email := os.Getenv("SEED_EMAIL")
    username := os.Getenv("SEED_USERNAME")
    password := os.Getenv("SEED_PASSWORD")

    if email == "" || username == "" || password == "" {
        return nil
    }

    hash, err := security.GeneratePasswordHash(password)
    if err != nil {
        return err
    }

    u := &models.User{
        Email:        email,
        Username:     username,
        PasswordHash: hash,
    }

    // Create if not exists
    return db.FirstOrCreate(u, "email = ?", email).Error
}
