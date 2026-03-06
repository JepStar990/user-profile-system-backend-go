package controllers

import (
    "crypto/sha256"
    "encoding/hex"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type AuthController struct{}

func (AuthController) Register(c *fiber.Ctx) error {
    var body dto.RegisterRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }

    if err := utils.Validate(body); err != nil {
        return err
    }

    user, err := services.Register(body)
    if err != nil {
        return fiber.NewError(fiber.StatusConflict, "Email or username already exists")
    }

    return c.JSON(fiber.Map{
        "message": "User registered successfully",
        "user_id": user.ID,
    })
}

func (AuthController) Login(c *fiber.Ctx) error {
    var body dto.LoginRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }

    user, err := services.Login(body)
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
    }

    access, refresh, err := services.GenerateTokenPair(user.ID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(dto.AuthResponse{
        AccessToken:  access,
        RefreshToken: refresh,
    })
}

func (AuthController) Refresh(c *fiber.Ctx) error {
    refresh := c.Get("X-Refresh-Token")
    if refresh == "" {
        return fiber.NewError(fiber.StatusBadRequest, "Missing refresh token header")
    }

    tokenHash := sha256.Sum256([]byte(refresh))
    hashHex := hex.EncodeToString(tokenHash[:])

    userID := c.Query("user_id")
    uid, err := uuid.Parse(userID)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid user")
    }

    valid, _ := repositories.IsTokenValid(uid, hashHex)
    if !valid {
        return fiber.NewError(fiber.StatusUnauthorized, "Refresh token invalid or expired")
    }

    access, newRefresh, err := services.GenerateTokenPair(uid)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(dto.AuthResponse{
        AccessToken:  access,
        RefreshToken: newRefresh,
    })
}

func (AuthController) Logout(c *fiber.Ctx) error {
    refresh := c.Get("X-Refresh-Token")
    if refresh == "" {
        return fiber.NewError(fiber.StatusBadRequest, "Missing refresh token header")
    }

    tokenHash := sha256.Sum256([]byte(refresh))
    hashHex := hex.EncodeToString(tokenHash[:])

    repositories.RevokeToken(hashHex)

    return c.JSON(fiber.Map{"message": "Logged out"})
}
