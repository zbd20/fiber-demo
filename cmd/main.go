package main

import (
	"flag"
	"log"

	_ "go.uber.org/automaxprocs"

	"github.com/zbd20/fiber-demo/internal/conf"
	"github.com/zbd20/fiber-demo/internal/router"
)

var cfgPath = flag.String("config", "configs/config.dev.yaml", "config path")

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	flag.Parse()

	if err := conf.InitConfig(*cfgPath); err != nil {
		log.Fatalf("init config error: %v", err)
	}

	app, err := router.InitRouter()
	if err != nil {
		log.Fatalf("new app error: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("start app error: %v", err)
	}
}
