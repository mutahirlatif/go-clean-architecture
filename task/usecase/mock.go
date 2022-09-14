package usecase

import (
	"context"
	"time"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"github.com/stretchr/testify/mock"
)

type TaskUseCaseMock struct {
	mock.Mock
}

func (m TaskUseCaseMock) CreateTask(ctx context.Context, user *models.User, taskDetails string, dueDate time.Time) error {
	args := m.Called(user, taskDetails, dueDate)

	return args.Error(0)
}

func (m TaskUseCaseMock) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	args := m.Called(user)

	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m TaskUseCaseMock) DeleteTask(ctx context.Context, user *models.User, id string) error {
	args := m.Called(user, id)

	return args.Error(0)
}
