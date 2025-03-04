package service

import (
	"context"
	"scheduler/internal/model"
	"sync"
	"time"
)

type Service interface {
	Create(ctx context.Context, description string, executeAt time.Time) (*model.Job, error)
	Cancel(jobID string) error
	GetAll() ([]*model.Job, error)
	GetByID(jobID string) (*model.Job, error)
	RunJob(ctx context.Context, jobID string) (*model.Job, error)
}

type service struct {
	mu   sync.RWMutex
	jobs map[string]*model.Job
}

func NewService() Service {
	return &service{
		jobs: make(map[string]*model.Job),
	}
}
