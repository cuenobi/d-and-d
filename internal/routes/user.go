package routes

import (
	"log/slog"

	"d-and-d/internal/handler"
	"d-and-d/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouteUserHandler(f *fiber.App, userService port.UserService, jwt port.JWT, validator *validator.Validate, logger *slog.Logger) *handler.UserHandler {
	handler := &handler.UserHandler{
		Fiber:       f,
		UserService: userService,
		JWT:         jwt,
		Validator:   validator,
		Logger:      logger,
	}

	f.Post("/login", handler.Login)
	user := f.Group("/user")
	user.Post("/register", handler.RegisterHandler)
	user.Get("/quests", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.GetAllQuest(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Get("/characters", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.GetAllCharacter(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Get("/races", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.GetAllRace(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Get("/classes", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.GetAllClass(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Get("/diff-lv", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.GetAllDifficultyLevel(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Post("/quest", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.CreateQuest(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Post("/character", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.CreateCharacter(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Put("/quest/:quest_id", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.UpdateQuest(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Put("/character/:character_id", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.UpdateCharacter(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Delete("/quest/:quest_id", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.DeleteQuest(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})
	user.Delete("/character/:character_id", jwt.Validate, func(ctx *fiber.Ctx) error {
		response := handler.DeleteCharacter(ctx)
		return ctx.Status(response.StatusCode).JSON(response)
	})

	return handler
}
