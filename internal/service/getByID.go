package service

import (
	"scheduler/internal/model"
)

func (s *service) GetByID(jobID string) (*model.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	job, exists := s.jobs[jobID]
	if !exists {
		return nil, model.ErrJobNotFound
	}

	return job, nil
}
