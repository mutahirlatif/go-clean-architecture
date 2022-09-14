package task

import (
	"context"

	"github.com/mutahirlatif/go-clean-architecture/models"
)

type Repository interface {
	CreateTask(ctx context.Context, user *models.User, bm *models.Task) error
	GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error)
	GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error)
	DeleteTask(ctx context.Context, user *models.User, id string) error
	UpdateTask(ctx context.Context, user *models.User, tm *models.Task, id string) error
}
