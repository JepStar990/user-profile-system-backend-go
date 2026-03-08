package utils

import "github.com/gofiber/fiber/v2"

// Common internal error codes used by ErrorHandler responses.
const (
    ErrBadRequest          = "BAD_REQUEST"
    ErrUnauthorized        = "UNAUTHORIZED"
    ErrForbidden           = "FORBIDDEN"
    ErrNotFound            = "NOT_FOUND"
    ErrConflict            = "CONFLICT"
    ErrUnprocessableEntity = "UNPROCESSABLE_ENTITY"
    ErrInternal            = "INTERNAL_SERVER_ERROR"
)

// NewHTTPError is a small helper to return Fiber errors consistently.
func NewHTTPError(status int, message string) error {
    return fiber.NewError(status, message)
}
