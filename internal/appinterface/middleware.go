package appinterface

import (
	"d-and-d/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerMiddleware interface {
	On(fn func(*fiber.Ctx) dto.HandlerResponse) fiber.Handler
}
