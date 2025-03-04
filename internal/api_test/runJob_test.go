package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"scheduler/internal/api"
	"scheduler/internal/model"
	service "scheduler/internal/service_test"
	"testing"
)

func TestAPI_RunJob(t *testing.T) {
	app := fiber.New()
	mockSvc := &service.MockJobService{}
	api.Register(app, mockSvc)
	ctx := mock.Anything

	t.Run("Run job - success", func(t *testing.T) {
		jobID := "1"
		mockSvc.On("RunJob", ctx, jobID).Return(&model.Job{ID: "1", Description: "Job 1"}, nil)

		req, err := http.NewRequest(fiber.MethodPost, "/jobs/1/run", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var job model.Job
		err = json.NewDecoder(resp.Body).Decode(&job)
		assert.NoError(t, err)
		assert.Equal(t, "Job 1", job.Description)
	})

	t.Run("Run job - error", func(t *testing.T) {
		jobID := "2"
		mockSvc.On("RunJob", ctx, jobID).Return(&model.Job{
			Status: model.StatusExecuted,
		}, model.ErrJobAlreadyExecuted)

		req, err := http.NewRequest(fiber.MethodPost, "/jobs/2/run", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Run job - error not found", func(t *testing.T) {
		jobID := "3"
		mockSvc.On("RunJob", ctx, jobID).Return(&model.Job{}, model.ErrJobNotFound)

		req, err := http.NewRequest(fiber.MethodPost, "/jobs/3/run", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}
