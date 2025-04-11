package allgood

import (
	"github.com/gofiber/fiber/v2"
)

// Middleware every response to have status 200 OK,
// regardless of what the handler tries to return.
// This ensures that the response is always a 200 OK.
func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		c.Status(fiber.StatusOK)
		return err
	}
}
