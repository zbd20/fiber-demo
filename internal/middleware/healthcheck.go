package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	LivenessPath  string
	ReadinessPath string
}

func HealthCheck(cfg *Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		if path != cfg.LivenessPath && path != cfg.ReadinessPath {
			return ctx.Next()
		}

		switch path {
		case cfg.LivenessPath:
			return ctx.Status(fiber.StatusOK).SendString("I am alive")
		case cfg.ReadinessPath:
			return ctx.Status(fiber.StatusOK).SendString("I am ready")
		}

		return nil
	}
}
