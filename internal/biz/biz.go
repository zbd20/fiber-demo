package biz

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"

	"github.com/zbd20/fiber-demo/internal/services"
)

var ControllerSet = wire.NewSet(NewBaseController)

type BaseController struct {
	bs  *services.BaseService
	fg  fiber.Router
	App *fiber.App
}

func NewBaseController(app *fiber.App, bs *services.BaseService) *BaseController {
	rg := app.Group("/demo/api/v1")
	app.Group("/").All("", adaptor.HTTPHandler(healthcheck.NewHandler()))


	bc := &BaseController{
		bs:  bs,
		fg:  rg,
		App: app,
	}

	// Controllers
	newDemoController(bc)

	return bc
}
