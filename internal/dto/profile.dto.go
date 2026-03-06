package dto

type UpdateProfileRequest struct {
    FullName string `json:"full_name" validate:"required,min=2,max=255"`
    Bio      string `json:"bio" validate:"max=2000"`
}

type ChangePasswordRequest struct {
    OldPassword string `json:"old_password" validate:"required"`
    NewPassword string `json:"new_password" validate:"required,min=8"`
}

type ProfileResponse struct {
    UserID    string `json:"user_id"`
    Email     string `json:"email"`
    Username  string `json:"username"`
    FullName  string `json:"full_name"`
    AvatarURL string `json:"avatar_url"`
    Bio       string `json:"bio"`
}
