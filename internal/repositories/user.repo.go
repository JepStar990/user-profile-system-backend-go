package repositories

import (
    "user-profile-system-backend-go/internal/db"
    "user-profile-system-backend-go/internal/models"

    "github.com/google/uuid"
)

func CreateUser(u *models.User) error {
    return db.DB.Create(u).Error
}

func FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := db.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}

func FindUserByID(id uuid.UUID) (*models.User, error) {
    var user models.User
    err := db.DB.First(&user, "id = ?", id).Error
    return &user, err
}
