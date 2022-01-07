package biz

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zbd20/fiber-demo/internal/models"
)

type demoController struct {
	*BaseController
}

func newDemoController(bc *BaseController) *demoController {
	dc := &demoController{bc}

	dc.fg.Post("/login", dc.login)
	return dc
}

func (dc *demoController) login(ctx *fiber.Ctx) error {
	var user models.Login
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	result, err := dc.bs.DemoService.Login(user)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString(fmt.Sprintf("hello %s", result.Username))
}
