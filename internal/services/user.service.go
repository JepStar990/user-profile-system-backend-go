package services

import (
    "errors"
    "mime/multipart"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/security"
    "user-profile-system-backend-go/internal/storage"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// -----------------------
//  PROFILE SERVICE
// -----------------------

func GetUserProfile(userID uuid.UUID) (*dto.ProfileResponse, error) {
    user, err := repositories.FindUserByID(userID)
    if err != nil {
        return nil, gorm.ErrRecordNotFound
    }

    profile, err := repositories.EnsureProfile(userID)
    if err != nil {
        return nil, err
    }

    return &dto.ProfileResponse{
        UserID:    user.ID.String(),
        Email:     user.Email,
        Username:  user.Username,
        FullName:  profile.FullName,
        AvatarURL: profile.AvatarURL,
        Bio:       profile.Bio,
    }, nil
}

func UpdateUserProfile(userID uuid.UUID, req dto.UpdateProfileRequest, c *fiber.Ctx) error {
    profile, err := repositories.EnsureProfile(userID)
    if err != nil {
        return err
    }

    profile.FullName = req.FullName
    profile.Bio = req.Bio
    err = repositories.SaveProfile(profile)

    if err == nil {
        LogActivity(userID, "update_profile", map[string]any{
            "full_name": req.FullName,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}

func ChangeUserPassword(userID uuid.UUID, req dto.ChangePasswordRequest, c *fiber.Ctx) error {
    user, err := repositories.FindUserByID(userID)
    if err != nil {
        return gorm.ErrRecordNotFound
    }

    valid, _ := security.ComparePasswordHash(req.OldPassword, user.PasswordHash)
    if !valid {
        return errors.New("incorrect old password")
    }

    newHash, err := security.GeneratePasswordHash(req.NewPassword)
    if err != nil {
        return err
    }

    user.PasswordHash = newHash
    err = repositories.CreateUser(user)

    if err == nil {
        LogActivity(userID, "change_password", nil, c.IP(), string(c.Context().UserAgent()))
    }

    return err
}

func UploadAvatar(userID uuid.UUID, file *multipart.FileHeader, c *fiber.Ctx) (string, error) {
    uploader := storage.NewS3Uploader()
    url, err := uploader.UploadAvatar(file, userID.String())
    if err != nil {
        return "", err
    }

    err = repositories.UpdateAvatarURL(userID, url)

    if err == nil {
        LogActivity(userID, "upload_avatar", map[string]any{
            "avatar_url": url,
        }, c.IP(), string(c.Context().UserAgent()))
    }

    return url, err
}
