package gateways

import (
	"fmt"
	"task_management_system/src/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) GetAllTasks(c *fiber.Ctx) error {
	return c.Status(200).JSON(entities.ResponseModel{Message: "Get All Task Success"})
}

func (h *HTTPGateway) AddNewTask(c *fiber.Ctx) error {
	payload := &entities.TaskModel{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(entities.ResponseModel{Message: fmt.Sprintf("Payload Invalid: %s", err.Error())})
	}
	if err := h.TaskService.AddNewTask(payload); err != nil {
		return err
	}
	return c.Status(200).JSON(entities.ResponseModel{Message: "Add New Task Success"})
}

func (h *HTTPGateway) EditTask(c *fiber.Ctx) error {
	return c.Status(200).JSON(entities.ResponseModel{Message: "Edit Task Success"})
}

func (h *HTTPGateway) DeleteTask(c *fiber.Ctx) error {
	return c.Status(200).JSON(entities.ResponseModel{Message: "Delete Task Success"})
}
