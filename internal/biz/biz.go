package biz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zbd20/fiber-demo/internal/services"
	"gorm.io/gorm"
)

type BaseController struct {
	db *gorm.DB
	bs *services.BaseService
	fg fiber.Router
}

func NewBaseController(app *fiber.App, db *gorm.DB) *BaseController {
	rg := app.Group("/demo/api/v1")

	bc := &BaseController{
		db: db,
		bs: services.NewBaseService(db),
		fg: rg,
	}

	newDemoController(bc)

	return bc
}
