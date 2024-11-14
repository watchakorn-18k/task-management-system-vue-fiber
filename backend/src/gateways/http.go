package gateways

import (
	service "task_management_system/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	TaskService service.ITasksService
}

func NewHTTPGateway(app *fiber.App, taskService service.ITasksService) {
	gateway := &HTTPGateway{
		TaskService: taskService,
	}

	RouteTask(*gateway, app)
}
