package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyubrr/strava-activity-editor/controller"
	"github.com/wahyubrr/strava-activity-editor/service"
)

func main() {
	app := fiber.New()

	stravaService := service.NewStravaService(os.Getenv("STRAVA_API"))
	controller := controller.NewController(stravaService)

	controller.NewRoutes(app)

	app.Listen(":8000")
}
