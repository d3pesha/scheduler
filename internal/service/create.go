package service

import (
	"context"
	"github.com/google/uuid"
	"scheduler/internal/model"
	"strings"
	"time"
)

func (s *service) Create(ctx context.Context, description string, executeAt time.Time) (*model.Job, error) {
	if len(strings.TrimSpace(description)) == 0 {
		return nil, model.ErrEmptyDescription
	}

	jobID := uuid.New().String()
	job := &model.Job{
		ID:          jobID,
		Description: description,
		ExecuteAt:   executeAt.UTC(),
		Status:      model.StatusScheduled,
		IsForcedRun: false,
	}

	if err := job.Validate(); err != nil {
		return nil, err
	}

	taskCtx, cancel := context.WithCancel(ctx)
	job.CancelFunc = cancel

	s.mu.Lock()
	s.jobs[jobID] = job
	s.mu.Unlock()

	go s.jobExecution(taskCtx, job)

	return job, nil
}
