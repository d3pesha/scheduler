package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"testing"
	"time"
)

func TestService_GetByID(t *testing.T) {
	svc := service.NewService()
	ctx := context.Background()

	t.Run("Valid job retrieval", func(t *testing.T) {
		job, err := svc.Create(ctx, "Test Job", time.Now().Add(5*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, job)

		retrievedJob, err := svc.GetByID(job.ID)
		assert.NoError(t, err)
		assert.Equal(t, job.ID, retrievedJob.ID)
	})

	t.Run("Retrieving non-existent job", func(t *testing.T) {
		_, err := svc.GetByID("nonexistent-job-id")
		assert.Error(t, err)
		assert.Equal(t, model.ErrJobNotFound, err)
	})
}
