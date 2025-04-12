package transport

import (
	_ "Obsidian/docs"
	handls "Obsidian/internal/handlers"
	"Obsidian/internal/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Transport struct {
	App *fiber.App
}

func NewTransport(repo *repository.TaskRepository) *Transport {
	app := fiber.New()

	registerRoutes(app, repo)

	return &Transport{
		App: app,
	}
}

func registerRoutes(app *fiber.App, repo *repository.TaskRepository) {
	taskHandler := handls.NewTasksHandler(repo)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/tasks", taskHandler.GetTasks)
	app.Post("/tasks", taskHandler.CreateTask)
	app.Put("/tasks/:id", taskHandler.UpdateTask)
	app.Delete("/tasks/:id", taskHandler.DeleteTask)

	routes := app.Stack()
	for _, routeList := range routes {
		for _, route := range routeList {
			fmt.Println(route.Name, route.Path, route.Method)
		}
	}
}
