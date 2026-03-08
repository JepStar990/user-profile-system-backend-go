package http

import "github.com/gofiber/fiber/v2"

// SuccessResponse is the unified success envelope.
// We keep this lightweight and consistent for frontend integration.
type SuccessResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Message string      `json:"message,omitempty"`
}

// OK returns a standard success response with optional data.
func OK(c *fiber.Ctx, data interface{}) error {
    return c.JSON(SuccessResponse{
        Success: true,
        Data:    data,
    })
}

// Message returns a standard success response with a message.
func Message(c *fiber.Ctx, msg string) error {
    return c.JSON(SuccessResponse{
        Success: true,
        Message: msg,
    })
}
