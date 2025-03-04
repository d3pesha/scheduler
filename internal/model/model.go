package model

import (
	"errors"
	"time"
)

const (
	StatusScheduled = "scheduled"
	StatusExecuting = "executing"
	StatusExecuted  = "executed"
	StatusCancelled = "cancelled"
)

var (
	ErrEmptyDescription    = errors.New("description is empty")
	ErrInvalidTime         = errors.New("executeAt can not be in the past")
	ErrJobNotFound         = errors.New("job not found")
	ErrJobAlreadyExecuted  = errors.New("job already executed")
	ErrJobAlreadyDone      = errors.New("job is already executed")
	ErrJobAlreadyCancelled = errors.New("job is already cancelled")
)

type Job struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	ExecuteAt   time.Time  `json:"executeAt"`
	ExecutedAt  *time.Time `json:"executedAt,omitempty"`
	Status      string     `json:"status"`
	IsForcedRun bool       `json:"-"`
	CancelFunc  func()     `json:"-"`
}

func (j *Job) Validate() error {
	if j.ExecuteAt.Before(time.Now()) {
		return ErrInvalidTime
	}

	return nil
}
