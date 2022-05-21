package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/go-pac/kages/model"
)

const (
	TokenKey        = "token"
	UserKey         = "user"
)

// todo: base64 decode ID
func GetUserFromCtx(ctx *fiber.Ctx) *model.User {
	if userData := ctx.Locals(UserKey); userData != nil {
		if user, ok := userData.(*model.User); ok {
			return user
		}
	}

	return nil
}

func GetTokenFromCtx(ctx *fiber.Ctx) string {
	if tokenVal := ctx.Locals(TokenKey); tokenVal != nil {
		if token, ok := tokenVal.(string); ok {
			return token
		}
	}

	return ""
}
