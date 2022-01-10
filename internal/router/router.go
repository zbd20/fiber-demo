package router

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"

	"github.com/zbd20/fiber-demo/internal/biz"
	"github.com/zbd20/fiber-demo/internal/conf"
	"github.com/zbd20/fiber-demo/internal/middleware"
)

var swagHandler fiber.Handler

var (
	AppSet   = wire.NewSet(NewFiber)
	RouteSet = wire.NewSet(NewRouter)
)

type Router struct {
	app *fiber.App
}

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.Page())
	app.Use(middleware.HealthCheck(&middleware.Config{
		LivenessPath:  "/live",
		ReadinessPath: "/ready",
	}))

	if swagHandler != nil {
		app.Group("/swagger").Get("/*", swagHandler)
	}

	return app
}

func NewRouter(bc *biz.BaseController) *Router {
	return &Router{
		app: bc.App,
	}
}

func (r *Router) Run() error {
	addr := conf.GetConfig().Addr
	log.Printf("start http server: %s", addr)

	return r.app.Listen(fmt.Sprintf(":%s", addr))
}
