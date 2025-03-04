package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"testing"
	"time"
)

func TestService_GetAll(t *testing.T) {
	svc := service.NewService()
	ctx := context.Background()

	t.Run("Get all jobs when no jobs exist", func(t *testing.T) {
		jobs, err := svc.GetAll()
		assert.Empty(t, jobs)
		assert.Equal(t, model.ErrJobNotFound, err)
	})

	t.Run("Get all jobs after creating jobs", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			job, err := svc.Create(ctx, "Test Job "+uuid.New().String(), time.Now().Add(time.Duration(i+1)*time.Second))
			assert.NoError(t, err)
			assert.NotNil(t, job)
		}

		jobs, err := svc.GetAll()
		assert.NoError(t, err)
		assert.Len(t, jobs, 3)
	})
}
