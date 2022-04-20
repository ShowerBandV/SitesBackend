package interpretor

import (
	"SitesBackend/tools"
	"github.com/gofiber/fiber/v2"
)

func MyInterpretor(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	//if token == "login" {
	//	return ctx.Next()
	//}
	if token == "" {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(tools.ResultSet{
			Data:   "NOTOKEN",
			Msg:    "please login before using",
			Status: 403,
		})
	}

	authoried, err := tools.ParseToken(token)
	if err != nil {
		return ctx.JSON(tools.ResultSet{
			Data:   "TOKENERROR",
			Msg:    "wrong token",
			Status: fiber.StatusForbidden,
		})
	}

	if !authoried {
		return ctx.JSON(tools.ResultSet{
			Data:   "TOKENPARSEERROR",
			Msg:    "token parse error",
			Status: fiber.StatusForbidden,
		})
	}
	return ctx.Next()
}
