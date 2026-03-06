package services

import (
    "errors"

    "user-profile-system-backend-go/internal/dto"
    "user-profile-system-backend-go/internal/models"
    "user-profile-system-backend-go/internal/repositories"
    "user-profile-system-backend-go/internal/security"
    "user-profile-system-backend-go/internal/storage"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

func GetUserProfile(userID uuid.UUID) (*dto.ProfileResponse, error) {
    user, err := repositories.FindUserByID(userID)
    if err != nil {
        return nil, gorm.ErrRecordNotFound
    }

    profile, err := repositories.GetProfileByUserID(userID)
    if err == gorm.ErrRecordNotFound {
        // create profile if doesn't exist
        profile = &models.UserProfile{
            UserID:   userID,
            FullName: "",
            AvatarURL: "",
            Bio:       "",
        }
        repositories.CreateOrUpdateProfile(profile)
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

func UpdateProfile(userID uuid.UUID, req dto.UpdateProfileRequest) error {
    profile, err := repositories.GetProfileByUserID(userID)
    if err == gorm.ErrRecordNotFound {
        profile = &models.UserProfile{
            UserID: userID,
        }
    }

    profile.FullName = req.FullName
    profile.Bio = req.Bio

    return repositories.CreateOrUpdateProfile(profile)
}

func ChangePassword(userID uuid.UUID, req dto.ChangePasswordRequest) error {
    user, err := repositories.FindUserByID(userID)
    if err != nil {
        return err
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
    return repositories.CreateUser(user)
}

func UploadAvatar(userID uuid.UUID, fileHeader *multipart.FileHeader) (string, error) {
    uploader := storage.NewS3Uploader()

    url, err := uploader.UploadAvatar(fileHeader, userID.String())
    if err != nil {
        return "", err
    }

    err = repositories.UpdateAvatar(userID, url)
    return url, err
}
