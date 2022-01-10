//go:generate wire
// +build wireinject

package router

import (
	"github.com/google/wire"

	"github.com/zbd20/fiber-demo/internal/biz"
	"github.com/zbd20/fiber-demo/internal/conf"
	"github.com/zbd20/fiber-demo/internal/db"
	"github.com/zbd20/fiber-demo/internal/services"
)

func InitRouter() (*Router, error) {
	wire.Build(RouteSet, AppSet, conf.ConfigSet, biz.ControllerSet, db.DatabaseSet, services.ServiceSet)
	return &Router{}, nil
}
