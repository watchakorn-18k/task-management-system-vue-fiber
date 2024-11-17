package main

import (
	"fmt"
	"os"
	"task_management_system/src/configuration"
	ds "task_management_system/src/domain/datasources"
	repo "task_management_system/src/domain/repositories"
	"task_management_system/src/gateways"
	"task_management_system/src/middlewares"
	sv "task_management_system/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/watchakorn-18k/scalar-go"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use("/api/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./src/docs/swagger.yaml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Task Management System  API",
			},
			DarkMode: true,
		})

		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(htmlContent)
	})
	app.Use(recover.New())
	app.Use(cors.New())

	mongoDB := ds.NewMongoDB(10)
	redisDB := ds.NewRedisConnection()

	tasksRepo := repo.NewTaskRepository(mongoDB)
	usersRepo := repo.NewUsersRepository(mongoDB)

	redisRepo := repo.NewRedisRepository(redisDB)

	taskSV := sv.NewTasksService(tasksRepo, redisRepo)
	usersSV := sv.NewUsersService(usersRepo)

	gateways.NewHTTPGateway(app, taskSV, usersSV)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
