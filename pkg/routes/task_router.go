package routes

import (
	"github.com/euferreira/pkg/handlers"
	"github.com/euferreira/pkg/repositories"
	"github.com/gofiber/fiber/v2"
)

func tasksRouter(app *fiber.App) *fiber.App {
	taskRepository := repositories.NewRepository()

	app.Get("/tasks", handlers.GetTasksHandler(taskRepository))
	app.Get("/tasks/:id", handlers.GetTaskHandler(taskRepository))
	app.Post("/tasks", handlers.CreateTaskHandler(taskRepository))
	app.Put("/tasks/:id", handlers.UpdateTaskHandler(taskRepository))
	app.Delete("/tasks/:id", handlers.DeleteTaskHandler(taskRepository))

	return app
}
