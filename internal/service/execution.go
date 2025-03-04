package service

import (
	"context"
	"scheduler/internal/model"
	"time"
)

func (s *service) jobExecution(ctx context.Context, job *model.Job) {
	if !job.IsForcedRun {
		timer := time.NewTimer(time.Until(job.ExecuteAt))
		defer timer.Stop()

		select {
		case <-timer.C:
		case <-ctx.Done():
			s.mu.Lock()
			job.Status = model.StatusCancelled
			s.mu.Unlock()
			return
		}
	}

	s.mu.Lock()
	if job.Status == model.StatusCancelled {
		s.mu.Unlock()
		return
	}

	now := time.Now().UTC()
	job.ExecutedAt = &now
	job.Status = model.StatusExecuting
	s.mu.Unlock()

	execTimer := time.NewTimer(10 * time.Second)
	defer execTimer.Stop()

	select {
	case <-execTimer.C:
		if job.Status != model.StatusCancelled {
			s.mu.Lock()
			job.Status = model.StatusExecuted
			s.mu.Unlock()
		}
	case <-ctx.Done():
		s.mu.Lock()
		job.Status = model.StatusCancelled
		s.mu.Unlock()
	}
}
