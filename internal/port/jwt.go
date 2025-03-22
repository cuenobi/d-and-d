package port

import "github.com/gofiber/fiber/v2"

type JWT interface {
	Generate(username string, role uint) string
	Validate(c *fiber.Ctx) error
	ValidateAdmin(c *fiber.Ctx) error
}
