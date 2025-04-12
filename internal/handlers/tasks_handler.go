package handlers

import (
	"Obsidian/internal/models"
	"Obsidian/internal/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type TasksHandler struct {
	Repo *repository.TaskRepository
}

func NewTasksHandler(repo *repository.TaskRepository) *TasksHandler {
	return &TasksHandler{
		Repo: repo,
	}
}

// CreateTask godoc
// @Summary Create a new task
// @Tags tasks
// @Accept */*
// @Produce json
// @Param task body models.TaskDto true "Task to create"
// @Success 201 {object} models.Task "Creates"
// @Failure 500 {object} models.ErrorResponse "Internal error"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Router /tasks [post]
func (t *TasksHandler) CreateTask(c *fiber.Ctx) error {
	task := &models.Task{}
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: err.Error()})
	}

	err := checkStatus(task.Status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: err.Error()})
	}

	err, tResult := t.Repo.CreateTask(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(tResult)
}

// GetTasks godoc
// @Summary Get all tasks
// @Accept */*
// @Produce json
// @Tags tasks
// @Success 200 {array} models.Task "Succeed gets tasks"
// @Failure 500 {object} models.ErrorResponse "Internal error"
// @Router /tasks [get]
func (t *TasksHandler) GetTasks(c *fiber.Ctx) error {
	err, tResult := t.Repo.GetTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(tResult)
}

// UpdateTask godoc
// @Summary Update a task
// @Tags tasks
// @Accept */*
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.TaskDto true "Updated task"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.ErrorResponse "Invalid date"
// @Failure 500 {object} models.ErrorResponse "Internal error"
// @Router /tasks/{id} [put]
func (t *TasksHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := &models.Task{}

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: err.Error()})
	}

	err := checkStatus(task.Status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: err.Error()})
	}

	err, tResult := t.Repo.UpdateTask(id, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(tResult)
}

// DeleteTask godoc
// @Summary Delete a task
// @Accept */*
// @Produce json
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 204 "Task was deleted"
// @Failure 500 {object} models.ErrorResponse "Internal error"
// @Router /tasks/{id} [delete]
func (t *TasksHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	err := t.Repo.DeleteTask(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func checkStatus(status string) error {
	if status != "new" && status != "in_progress" && status != "done" {
		return fmt.Errorf("invalid status: %s", status)
	}
	return nil
}
