package routes

import (
	"log/slog"

	"d-and-d/internal/handler"
	"d-and-d/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouteAdminHandler(f *fiber.App, jwt port.JWT, validator *validator.Validate, adminService port.AdminService, logger *slog.Logger) *handler.AdminHandler {
	handler := &handler.AdminHandler{
		Fiber:        f,
		JWT:          jwt,
		Validator:    validator,
		AdminService: adminService,
		Logger:       logger,
	}

	admin := f.Group("/admin")
	admin.Post("/race", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.CreateRace(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Post("/class", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.CreateClass(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Post("/diff-lv", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.CreateDifficultyLevel(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Put("/race", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.UpdateRace(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Put("/class", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.UpdateClass(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Put("/diff-lv", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.UpdateDifficultyLevel(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Delete("/race/:race_id", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.DeleteRace(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Delete("/class/:class_id", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.DeleteClass(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	admin.Delete("/diff-lv/:difficulty_level_id", jwt.ValidateAdmin, func(ctx *fiber.Ctx) error {
		response := handler.DeleteDifficultyLevel(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})

	return handler
}
