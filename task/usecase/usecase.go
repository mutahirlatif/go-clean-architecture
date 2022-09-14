package usecase

import (
	"context"
	"time"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"github.com/mutahirlatif/go-clean-architecture/task"
)

type TaskUseCase struct {
	taskRepo task.Repository
}

func NewTaskUseCase(taskRepo task.Repository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (t TaskUseCase) CreateTask(ctx context.Context, user *models.User, taskDetails string, dueDate time.Time) error {
	tm := &models.Task{
		TaskDetail: taskDetails,
		DueDate:    dueDate,
	}

	return t.taskRepo.CreateTask(ctx, user, tm)
}

func (t TaskUseCase) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	return t.taskRepo.GetTasks(ctx, user)
}

func (t TaskUseCase) DeleteTask(ctx context.Context, user *models.User, id string) error {
	return t.taskRepo.DeleteTask(ctx, user, id)
}

func (t TaskUseCase) GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error) {
	return t.taskRepo.GetTaskByID(ctx, user, id)
}

func (t TaskUseCase) UpdateTask(ctx context.Context, user *models.User, taskDetails string, dueDate time.Time, id string) error {
	tm := &models.Task{
		TaskDetail: taskDetails,
		DueDate:    dueDate,
	}
	return t.taskRepo.UpdateTask(ctx, user, tm, id)
}
