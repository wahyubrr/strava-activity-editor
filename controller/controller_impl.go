package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyubrr/strava-activity-editor/service"
)

type ControllerImpl struct {
	Service service.StravaService
}

func NewController(service service.StravaService) Controller {
	return &ControllerImpl{
		Service: service,
	}
}

func (controller *ControllerImpl) NewRoutes(api fiber.Router) {
	api.Get("/activities", controller.GetAllActivities)
	api.Get("/activities/:activityId", controller.GetOneActivities)
	api.Get("/gear/:gearId", controller.GetOneGear)
}

func (controller *ControllerImpl) GetAllActivities(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	per_page, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	fmt.Println("controller: get activities", page, per_page)
	controller.Service.GetAllActivities(page, per_page)
	return nil
}

func (controller *ControllerImpl) GetOneActivities(c *fiber.Ctx) error {
	activityId := c.Params("activityId", "")
	if activityId == "" {
		return fiber.ErrBadRequest
	}

	controller.Service.GetOneActivities(activityId)
	return nil
}

func (controller *ControllerImpl) GetOneGear(c *fiber.Ctx) error {
	gearId := c.Params("gearId", "")
	if gearId == "" {
		return fiber.ErrBadRequest
	}

	controller.Service.GetOneGear(gearId)
	return nil
}
