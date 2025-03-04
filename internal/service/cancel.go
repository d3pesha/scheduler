package service

import (
	"scheduler/internal/model"
)

func (s *service) Cancel(jobID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	job, exists := s.jobs[jobID]
	if !exists {
		return model.ErrJobNotFound
	}

	if job.Status == model.StatusExecuted {
		return model.ErrJobAlreadyDone
	} else if job.Status == model.StatusCancelled {
		return model.ErrJobAlreadyCancelled
	}

	if job.CancelFunc != nil {
		job.CancelFunc()
	}

	job.Status = model.StatusCancelled

	return nil
}
