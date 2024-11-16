package gateways

import (
	service "task_management_system/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	TaskService  service.ITasksService
	UsersService service.IUsersService
}

func NewHTTPGateway(app *fiber.App, taskService service.ITasksService, usersService service.IUsersService) {
	gateway := &HTTPGateway{
		TaskService:  taskService,
		UsersService: usersService,
	}

	RouteTask(*gateway, app)
	RouteUsers(*gateway, app)
}
