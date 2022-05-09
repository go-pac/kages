package auth

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrMissing = &fiber.Error{
		Code:    401000,
		Message: "Missing token",
	}
	ErrInvalid = &fiber.Error{
		Code:    401001,
		Message: "Invalid token",
	}
)
