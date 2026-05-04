package dto

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email,max=255"`
    Username string `json:"username" validate:"required,min=3,max=50,alphanumunicode"`
    Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required,max=128"`
}

type AuthResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}
