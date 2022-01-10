// +build dev

package router

import (
	"github.com/arsmn/fiber-swagger/v2"
	_ "github.com/zbd20/fiber-demo/api/swagger"
)

func init() {
	swagHandler = swagger.New()
}
