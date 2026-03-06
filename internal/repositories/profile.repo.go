package repositories

import (
    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

func GetProfileByUserID(userID uuid.UUID) (*models.UserProfile, error) {
    var p models.UserProfile
    err := db.DB.Where("user_id = ?", userID).First(&p).Error
    return &p, err
}

func CreateProfile(p *models.UserProfile) error {
    return db.DB.Create(p).Error
}

func SaveProfile(p *models.UserProfile) error {
    return db.DB.Save(p).Error
}

func UpdateAvatarURL(userID uuid.UUID, url string) error {
    return db.DB.Model(&models.UserProfile{}).
        Where("user_id = ?", userID).
        Update("avatar_url", url).Error
}

func ProfileExists(userID uuid.UUID) bool {
    var count int64
    db.DB.Model(&models.UserProfile{}).Where("user_id = ?", userID).Count(&count)
    return count > 0
}

func EnsureProfile(userID uuid.UUID) (*models.UserProfile, error) {
    if !ProfileExists(userID) {
        p := &models.UserProfile{
            UserID: userID,
        }
        err := CreateProfile(p)
        return p, err
    }

    return GetProfileByUserID(userID)
}
