package router

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/zbd20/fiber-demo/internal/biz"
	"github.com/zbd20/fiber-demo/internal/conf"
	"github.com/zbd20/fiber-demo/internal/db"
	"github.com/zbd20/fiber-demo/internal/middleware"
)

var swagHandler fiber.Handler

type Router struct {
	app *fiber.App
	bc  *biz.BaseController
}

func NewRouter() *Router {
	serverConfig := conf.GetConfig()
	mdb, err := db.NewMySQLClient(serverConfig)
	if err != nil {
		log.Fatalf("new mysql client error: %v", err)
		return nil
	}

	app := fiber.New()

	app.Use(middleware.Page())
	app.Use(logger.New())

	bc := biz.NewBaseController(app, mdb)

	r := &Router{
		app: app,
		bc:  bc,
	}

	if swagHandler != nil {
		app.Get("/swagger/*", swagHandler)
	}
	return r
}

func (r *Router) Run() error {
	addr := conf.GetConfig().Addr
	log.Printf("start http server: %s", addr)

	return r.app.Listen(fmt.Sprintf(":%s", addr))
}
