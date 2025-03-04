package service

import (
	"scheduler/internal/model"
)

func (s *service) GetAll() ([]*model.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.jobs) == 0 {
		return nil, model.ErrJobNotFound
	}

	jobs := make([]*model.Job, 0, len(s.jobs))
	for _, job := range s.jobs {
		jobs = append(jobs, job)
	}

	return jobs, nil
}
