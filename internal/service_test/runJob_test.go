package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"testing"
	"time"
)

func TestService_RunJob(t *testing.T) {
	svc := service.NewService()
	ctx := context.Background()

	job, err := svc.Create(ctx, "Test Job", time.Now().Add(1*time.Hour))
	assert.NoError(t, err)
	assert.NotNil(t, job)

	t.Run("Successful execution", func(t *testing.T) {
		runJob, err := svc.RunJob(ctx, job.ID)
		assert.NoError(t, err)
		assert.Equal(t, model.StatusExecuting, runJob.Status)
	})

	t.Run("Non-existent job should return error", func(t *testing.T) {
		_, err = svc.RunJob(ctx, "nonexistent-job-id")
		assert.Error(t, err)
		assert.Equal(t, model.ErrJobNotFound, err)
	})
}
