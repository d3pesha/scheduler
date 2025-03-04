package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"scheduler/internal/model"
	"time"
)

type MockJobService struct {
	mock.Mock
}

func (m *MockJobService) Create(ctx context.Context, description string, executeAt time.Time) (*model.Job, error) {
	args := m.Called(ctx, description, executeAt)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJobService) Cancel(jobID string) error {
	args := m.Called(jobID)
	return args.Error(0)
}

func (m *MockJobService) GetAll() ([]*model.Job, error) {
	args := m.Called()
	return args.Get(0).([]*model.Job), args.Error(1)
}

func (m *MockJobService) GetByID(jobID string) (*model.Job, error) {
	args := m.Called(jobID)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJobService) RunJob(ctx context.Context, jobID string) (*model.Job, error) {
	args := m.Called(ctx, jobID)
	return args.Get(0).(*model.Job), args.Error(1)
}
