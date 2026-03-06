package controllers

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type ProfileController struct{}

func (ProfileController) GetProfile(c *fiber.Ctx) error {
    userIDStr := c.Locals("user_id").(string)
    userID, _ := uuid.Parse(userIDStr)

    profile, err := services.GetUserProfile(userID)
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Profile not found")
    }

    return c.JSON(profile)
}

func (ProfileController) UpdateProfile(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.UpdateProfileRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid payload")
    }

    if err := utils.Validate(body); err != nil {
        return err
    }

    if err := services.UpdateProfile(userID, body); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{"message": "Profile updated"})
}

func (ProfileController) ChangePassword(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    var body dto.ChangePasswordRequest
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid payload")
    }

    if err := utils.Validate(body); err != nil {
        return err
    }

    err := services.ChangePassword(userID, body)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    return c.JSON(fiber.Map{"message": "Password changed"})
}

func (ProfileController) UploadAvatar(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    file, err := c.FormFile("avatar")
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Avatar file required")
    }

    url, err := services.UploadAvatar(userID, file)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    return c.JSON(fiber.Map{
        "message":    "Avatar updated",
        "avatar_url": url,
    })
}
