package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"testing"
	"time"
)

func TestService_Cancel(t *testing.T) {
	svc := service.NewService()
	ctx := context.Background()

	t.Run("Cancel job successfully", func(t *testing.T) {
		job, err := svc.Create(ctx, "Test Job", time.Now().Add(5*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, job)

		err = svc.Cancel(job.ID)
		assert.NoError(t, err)

		retrievedJob, err := svc.GetByID(job.ID)
		assert.NoError(t, err)
		assert.Equal(t, model.StatusCancelled, retrievedJob.Status)
	})

	t.Run("Cancel already cancelled job", func(t *testing.T) {
		job, err := svc.Create(ctx, "Test Job", time.Now().Add(5*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, job)

		err = svc.Cancel(job.ID)
		assert.NoError(t, err)

		err = svc.Cancel(job.ID)
		assert.Error(t, err)
		assert.Equal(t, model.ErrJobAlreadyCancelled, err)
	})

	t.Run("Cancel non-existent job", func(t *testing.T) {
		err := svc.Cancel("nonexistent-job-id")
		assert.Error(t, err)
		assert.Equal(t, model.ErrJobNotFound, err)
	})
}
