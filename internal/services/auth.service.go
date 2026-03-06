package services

import (
    "crypto/sha256"
    "encoding/hex"
    "time"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/security"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

func Register(req dto.RegisterRequest) (*models.User, error) {
    hash, err := security.GeneratePasswordHash(req.Password)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:        req.Email,
        Username:     req.Username,
        PasswordHash: hash,
    }

    err = repositories.CreateUser(user)
    return user, err
}

func Login(req dto.LoginRequest) (*models.User, error) {
    user, err := repositories.FindUserByEmail(req.Email)
    if err == gorm.ErrRecordNotFound {
        return nil, err
    }

    ok, _ := security.ComparePasswordHash(req.Password, user.PasswordHash)
    if !ok {
        return nil, gorm.ErrRecordNotFound
    }

    return user, nil
}

func GenerateTokenPair(userID uuid.UUID) (string, string, error) {
    access, err := security.GenerateAccessToken(userID)
    if err != nil {
        return "", "", err
    }

    refresh, err := security.GenerateRefreshToken()
    if err != nil {
        return "", "", err
    }

    hash := sha256.Sum256([]byte(refresh))
    rt := &models.RefreshToken{
        UserID:    userID,
        TokenHash: hex.EncodeToString(hash[:]),
        ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
    }

    repositories.StoreRefreshToken(rt)

    return access, refresh, nil
}
