package utils

import (
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Validate(data interface{}) error {
    err := validate.Struct(data)
    if err == nil {
        return nil
    }

    return fiber.NewError(fiber.StatusBadRequest, err.Error())
}
