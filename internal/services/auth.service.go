package services

import (
    "crypto/sha256"
    "encoding/hex"
    "time"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/security"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

func Register(req dto.RegisterRequest, c *fiber.Ctx) (*models.User, error) {
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
    if err == nil {
        LogActivity(user.ID, "register", nil, c.IP(), string(c.Context().UserAgent()))
    }
    return user, err
}

func Login(req dto.LoginRequest, c *fiber.Ctx) (*models.User, error) {
    user, err := repositories.FindUserByEmail(req.Email)
    if err != nil {
        return nil, gorm.ErrRecordNotFound
    }

    ok, _ := security.ComparePasswordHash(req.Password, user.PasswordHash)
    if !ok {
        return nil, gorm.ErrRecordNotFound
    }

    LogActivity(user.ID, "login", nil, c.IP(), string(c.Context().UserAgent()))
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

func Logout(userID uuid.UUID, refresh string, c *fiber.Ctx) error {
    hash := sha256.Sum256([]byte(refresh))
    err := repositories.RevokeToken(hex.EncodeToString(hash[:]))

    if err == nil {
        LogActivity(userID, "
