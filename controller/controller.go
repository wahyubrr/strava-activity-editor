package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	NewRoutes(api fiber.Router)
	GetAllActivities(*fiber.Ctx) error
	GetOneActivities(c *fiber.Ctx) error
	GetOneGear(c *fiber.Ctx) error
}
