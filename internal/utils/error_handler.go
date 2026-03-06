package utils

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// ErrorResponse is the unified JSON error schema returned by the API.
type ErrorResponse struct {
    Success bool        `json:"success"`
    Error   ErrorObject `json:"error"`
}

type ErrorObject struct {
    Code      string      `json:"code"`
    Message   string      `json:"message"`
    RequestID string      `json:"request_id"`
    Details   interface{} `json:"details,omitempty"`
}

// ErrorHandler is the global Fiber error handler.
// It converts all panic and Fiber errors into a safe, clean JSON envelope.
func ErrorHandler(c *fiber.Ctx, err error) error {
    reqID := c.Locals("request_id")
    if reqID == nil {
        reqID = uuid.New().String()
    }

    // Map Fiber errors → HTTP status + code
    code := fiber.StatusInternalServerError
    errorCode := "INTERNAL_SERVER_ERROR"
    msg := "An unexpected error occurred"

    // Fiber's built‑in errors
    if fe, ok := err.(*fiber.Error); ok {
        code = fe.Code
        msg = fe.Message

        switch fe.Code {
        case 400:
            errorCode = "BAD_REQUEST"
        case 401:
            errorCode = "UNAUTHORIZED"
        case 403:
            errorCode = "FORBIDDEN"
        case 404:
            errorCode = "NOT_FOUND"
        case 409:
            errorCode = "CONFLICT"
        case 422:
            errorCode = "UNPROCESSABLE_ENTITY"
        }
    }

    res := ErrorResponse{
        Success: false,
        Error: ErrorObject{
            Code:      errorCode,
            Message:   msg,
            RequestID: reqID.(string),
        },
    }

    return c.Status(code).JSON(res)
}
