package gateways

import (
	"fmt"
	"task_management_system/src/domain/entities"
	"task_management_system/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) RegisterUser(c *fiber.Ctx) error {
	payload := &entities.UserModel{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(entities.ResponseModel{Message: fmt.Sprintf("Payload Invalid: %s", err.Error())})
	}
	if err := h.UsersService.RegisterUsers(payload); err != nil {
		return c.Status(400).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return c.Status(200).JSON(entities.ResponseModel{Message: "Register User Success"})
}

func (h *HTTPGateway) LoginUser(c *fiber.Ctx) error {
	payload := &entities.UserModel{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(entities.ResponseModel{Message: fmt.Sprintf("Payload Invalid: %s", err.Error())})
	}
	token, err := h.UsersService.Login(payload)
	if err != nil {
		return c.Status(400).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return c.Status(200).JSON(entities.ResponseModel{Message: "Login User Success", Data: token})
}

func (h *HTTPGateway) GetProfile(c *fiber.Ctx) error {
	token, err := middlewares.DecodeJWTToken(c)
	if err != nil {
		return c.Status(400).JSON(entities.ResponseModel{Message: err.Error()})
	}
	userID := token.UserID
	if userID == "" {
		return c.Status(400).JSON(entities.ResponseModel{Message: "User ID Not Found"})
	}
	user, err := h.UsersService.GetProfile(token.UserID)
	if err != nil {
		return c.Status(400).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return c.Status(200).JSON(entities.ResponseModel{Message: "Get Profile Success", Data: fiber.Map{"user_id": user.UserID, "username": user.Username}})
}
