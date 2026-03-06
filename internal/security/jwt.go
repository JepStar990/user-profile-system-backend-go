package security

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
)

type AccessClaims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateAccessToken(userID uuid.UUID) (string, error) {
    secret := []byte(os.Getenv("JWT_ACCESS_SECRET"))
    exp, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRES"))

    claims := AccessClaims{
        UserID: userID.String(),
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
        SignedString(secret)
}

func GenerateRefreshToken() (string, error) {
    secret := []byte(os.Getenv("JWT_REFRESH_SECRET"))
    exp, _ := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRES"))

    claims := jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }

    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
        SignedString(secret)
}
