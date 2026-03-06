package security

import (
    "os"

    "github.com/golang-jwt/jwt/v5"
    "github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
    auth := c.Get("Authorization")
    if auth == "" {
        return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
    }

    tokenStr := auth[len("Bearer "):]

    token, err := jwt.ParseWithClaims(tokenStr, &AccessClaims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
    })
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
    }

    claims := token.Claims.(*AccessClaims)
    c.Locals("user_id", claims.UserID)

    return c.Next()
}

