package mock

import (
	"context"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"github.com/stretchr/testify/mock"
)

type TaskStorageMock struct {
	mock.Mock
}

func (s *TaskStorageMock) CreateTask(ctx context.Context, user *models.User, tm *models.Task) error {
	args := s.Called(user, tm)

	return args.Error(0)
}

func (s *TaskStorageMock) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Task), args.Error(1)
}

func (s *TaskStorageMock) DeleteTask(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}

func (s *TaskStorageMock) GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error) {
	args := s.Called(user)

	return args.Get(0).(*models.Task), args.Error(1)
}

func (s *TaskStorageMock) UpdateTask(ctx context.Context, user *models.User, tm *models.Task, id string) error {
	args := s.Called(user, tm)

	return args.Error(0)
}
