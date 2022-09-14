package task

import (
	"context"
	"time"

	"github.com/mutahirlatif/go-clean-architecture/models"
)

type UseCase interface {
	CreateTask(ctx context.Context, user *models.User, taskDetails string, dueDate time.Time) error
	GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error)
	GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error)
	DeleteTask(ctx context.Context, user *models.User, id string) error
	UpdateTask(ctx context.Context, user *models.User, taskDetails string, dueDate time.Time, id string) error
}
