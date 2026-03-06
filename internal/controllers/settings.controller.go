package controllers

import (
    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/services"
    "user-profile-system-backend-go/internal/utils"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type SettingsController struct{}

// GET /settings  → return entire JSON object
func (SettingsController) GetSettings(c *fiber.Ctx) error {
    userID, _ := uuid.Parse(c.Locals("user_id").(string))

    settings, err := services.GetSettings(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(fiber.Map{
        "settings": settings,
    })
}

// Helper for PUT endpoints
func bindAndValidate[T any](c *fiber.Ctx, dst *T) error {
    if err := c.BodyParser(dst); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
    }
    if err := utils.Validate(*dst); err != nil {
        return err
    }
    return nil
}

// SECTION ROUTES

func (SettingsController) UpdateAudio(c *fiber.Ctx) error {
    var body dto.AudioSettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdateAudioSettings(userID, body)

    return c.JSON(fiber.Map{"message": "Audio settings updated"})
}

func (SettingsController) UpdateVoice(c *fiber.Ctx) error {
    var body dto.VoiceSettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdateVoiceSettings(userID, body)

    return c.JSON(fiber.Map{"message": "Voice settings updated"})
}

func (SettingsController) UpdateLive(c *fiber.Ctx) error {
    var body dto.LiveRadioSettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdateLiveRadioSettings(userID, body)

    return c.JSON(fiber.Map{"message": "Live radio settings updated"})
}

func (SettingsController) UpdateNotifications(c *fiber.Ctx) error {
    var body dto.NotificationSettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdateNotificationSettings(userID, body)

    return c.JSON(fiber.Map{"message": "Notification settings updated"})
}

func (SettingsController) UpdateAppearance(c *fiber.Ctx) error {
    var body dto.AppearanceSettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdateAppearanceSettings(userID, body)

    return c.JSON(fiber.Map{"message": "Appearance settings updated"})
}

func (SettingsController) UpdatePrivacy(c *fiber.Ctx) error {
    var body dto.PrivacySettingsRequest
    if err := bindAndValidate(c, &body); err != nil {
        return err
    }

    userID, _ := uuid.Parse(c.Locals("user_id").(string))
    services.UpdatePrivacySettings(userID, body)

    return c.JSON(fiber.Map{"message": "Privacy settings updated"})
}
