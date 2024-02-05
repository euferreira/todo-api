package handlers

import (
	"github.com/euferreira/pkg/entities"
	"github.com/euferreira/pkg/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetTasksHandler(taskRepository *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks := taskRepository.GetTasks()

		if len(tasks) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "No tasks found"})
		}

		return c.Status(200).JSON(fiber.Map{"tasks": tasks})
	}
}

func GetTaskHandler(taskRepository *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		task := taskRepository.GetTask(id)

		if task.ID == "" {
			return c.Status(404).JSON(fiber.Map{
				"message": "No task found",
				"id":      id,
			})
		}

		return c.Status(200).JSON(fiber.Map{"task": task})
	}
}

func CreateTaskHandler(taskRepository *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(entities.Task)

		if err := c.BodyParser(task); err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
		}

		taskRepository.CreateTask(*task)

		return c.Status(201).JSON(fiber.Map{"message": "Task created"})
	}
}

func UpdateTaskHandler(taskRepository *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		task := new(entities.Task)

		if err := c.BodyParser(task); err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
		}

		taskRepository.UpdateTask(id, *task)

		return c.Status(200).JSON(fiber.Map{"message": "Task updated"})
	}
}

func DeleteTaskHandler(taskRepository *repositories.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		taskRepository.DeleteTask(id)

		return c.Status(200).JSON(fiber.Map{"message": "Task deleted"})
	}
}
