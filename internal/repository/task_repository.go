package repository

import (
	"Obsidian/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type TaskRepository struct {
	repo *pgxpool.Pool
}

func NewTaskRepository(mainRepo *Repository) *TaskRepository {
	return &TaskRepository{
		repo: mainRepo.DB,
	}
}

func (t *TaskRepository) CreateTask(model *models.Task) (error, *models.Task) {
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := t.repo.QueryRow(
		context.Background(),
		query,
		model.Title, model.Description, model.Status,
	).Scan(&model.ID, &model.CreatedAt, &model.UpdatedAt)

	if err != nil {
		return err, nil
	}

	return nil, model
}

func (t *TaskRepository) GetTasks() (error, []models.Task) {
	rows, err := t.repo.Query(context.Background(), `SELECT id, title, description, status, created_at, updated_at FROM tasks`)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	var tasks = make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return err, nil
		}
		tasks = append(tasks, task)
	}

	return nil, tasks
}

func (t *TaskRepository) UpdateTask(id string, task *models.Task) (error, *models.Task) {
	query := `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=now() WHERE id=$4 RETURNING id, created_at, updated_at`
	err := t.repo.QueryRow(context.Background(), query, task.Title, task.Description, task.Status, id).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return err, nil
	}

	return nil, task
}

func (t *TaskRepository) DeleteTask(id string) error {
	_, err := t.repo.Exec(context.Background(), `DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
