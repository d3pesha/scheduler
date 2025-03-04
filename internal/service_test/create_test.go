package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"sync"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	svc := service.NewService()
	var wg sync.WaitGroup
	jobCount := 100

	t.Run("Create multiple jobs concurrently", func(t *testing.T) {
		for i := 0; i < jobCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				ctx := context.Background()
				job, err := svc.Create(ctx, "Job-"+uuid.New().String(), time.Now().Add(time.Duration(i+1)*time.Second))
				assert.NoError(t, err)
				assert.NotNil(t, job)
			}(i)
		}

		wg.Wait()

		jobs, err := svc.GetAll()
		assert.NoError(t, err)
		assert.Len(t, jobs, jobCount)
	})

	t.Run("Cancel jobs concurrently", func(t *testing.T) {
		jobs, err := svc.GetAll()
		assert.NoError(t, err)

		for _, job := range jobs {
			wg.Add(1)
			go func(jobID string) {
				defer wg.Done()

				err := svc.Cancel(jobID)
				assert.NoError(t, err)
			}(job.ID)
		}

		wg.Wait()

		for _, job := range jobs {
			retrievedJob, err := svc.GetByID(job.ID)
			assert.NoError(t, err)
			assert.Equal(t, model.StatusCancelled, retrievedJob.Status)
		}
	})
}
