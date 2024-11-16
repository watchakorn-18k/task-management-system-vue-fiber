package gateways

import (
	"task_management_system/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteTask(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/task")
	api.Get("/get_all", gateway.GetAllTasks)
	api.Post("/add_task", gateway.AddNewTask)
	api.Put("/edit_task/:task_id", gateway.EditTask)
	api.Delete("/delete_task/:task_id", gateway.DeleteTask)
}

func RouteUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/user")
	api.Post("/register", gateway.RegisterUser)
	api.Post("/login", gateway.LoginUser)
	api.Get("/get_profile", middlewares.SetJWtHeaderHandler(), gateway.GetProfile)
}
