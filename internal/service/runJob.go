package service

import (
	"context"
	"scheduler/internal/model"
	"time"
)

func (s *service) RunJob(ctx context.Context, jobID string) (*model.Job, error) {
	s.mu.Lock()

	job, exists := s.jobs[jobID]
	if !exists {
		s.mu.Unlock()
		return nil, model.ErrJobNotFound
	}

	if job.Status == model.StatusExecuted {
		s.mu.Unlock()
		return nil, model.ErrJobAlreadyDone
	} else if job.Status == model.StatusCancelled {
		s.mu.Unlock()
		return nil, model.ErrJobAlreadyCancelled
	}

	job.IsForcedRun = true
	job.Status = model.StatusExecuting

	ctxTask, cancel := context.WithTimeout(ctx, 15*time.Second)
	job.CancelFunc = cancel
	s.mu.Unlock()

	go s.jobExecution(ctxTask, job)

	return job, nil
}
