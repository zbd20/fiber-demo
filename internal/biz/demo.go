package biz

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zbd20/fiber-demo/internal/models"
	"github.com/zbd20/fiber-demo/internal/pkg"
)

type demoController struct {
	*BaseController
}

func newDemoController(bc *BaseController) *demoController {
	dc := &demoController{bc}

	dc.fg.Post("/login", dc.login)

	return dc
}

// @Summary Login
// @Description Login Controller
// @Tags Login
// @version 1.0
// @Accept json
// @Produce json
// @Success 200 {object} models.Login OK
// @Failure 400 {string} string ERROR
// @Router /demo/api/v1/login [post]
func (dc *demoController) login(ctx *fiber.Ctx) error {
	var user models.Login
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	result, err := dc.bs.DemoService.Login(user)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(models.NewResult(0, nil, 0, pkg.Success, result))
}
