package repositories

import (
    "time"
    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
)

func StoreRefreshToken(rt *models.RefreshToken) error {
    return db.DB.Create(rt).Error
}

func RevokeToken(tokenHash string) error {
    return db.DB.Model(&models.RefreshToken{}).
        Where("token_hash = ?", tokenHash).
        Update("revoked", true).Error
}

func IsTokenValid(userID uuid.UUID, tokenHash string) (bool, error) {
    var rt models.RefreshToken
    err := db.DB.Where(
        "user_id = ? AND token_hash = ? AND revoked = false AND expires_at > ?",
        userID, tokenHash, time.Now(),
    ).First(&rt).Error

    return err == nil, err
}
