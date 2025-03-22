package routes

import (
	"log/slog"

	"d-and-d/config"
	"d-and-d/internal/handler"
	"d-and-d/internal/port"

	"github.com/gofiber/fiber/v2"
)

func NewRoutePublicHandler(f *fiber.App, publicService port.PublicService, logger *slog.Logger, cfg *config.Config) *handler.PublicHandler {
	handler := &handler.PublicHandler{
		Fiber:         f,
		PublicService: publicService,
		Logger:        logger,
		Config:        cfg,
	}

	public := f.Group("/public")
	public.Get("/quests", func(ctx *fiber.Ctx) error {
		response := handler.GetPublicQuest(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	public.Get("/characters", func(ctx *fiber.Ctx) error {
		response := handler.GetPublicCharacter(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	public.Get("/images", handler.GetImageFile)

	return handler
}
