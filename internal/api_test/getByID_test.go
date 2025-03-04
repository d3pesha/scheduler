package api

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"scheduler/internal/api"
	"scheduler/internal/model"
	service "scheduler/internal/service_test"
	"testing"
)

func TestAPI_GetJobByID(t *testing.T) {
	app := fiber.New()
	mockSvc := &service.MockJobService{}
	api.Register(app, mockSvc)

	t.Run("Get job by ID - success", func(t *testing.T) {
		jobID := "1"
		mockSvc.On("GetByID", jobID).Return(&model.Job{ID: "1", Description: "Job 1"}, nil)

		req, err := http.NewRequest(fiber.MethodGet, "/jobs/1", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var job model.Job
		err = json.NewDecoder(resp.Body).Decode(&job)
		assert.NoError(t, err)
		assert.Equal(t, "Job 1", job.Description)
	})

	t.Run("Get job by ID - not found", func(t *testing.T) {
		jobID := "2"
		mockSvc.On("GetByID", jobID).Return(&model.Job{}, model.ErrJobNotFound)

		req, err := http.NewRequest(fiber.MethodGet, "/jobs/2", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	t.Run("Get job by ID - error", func(t *testing.T) {
		jobID := "3"
		mockSvc.On("GetByID", jobID).Return(&model.Job{}, errors.New("some error"))

		req, err := http.NewRequest(fiber.MethodGet, "/jobs/3", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
