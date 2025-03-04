package api_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"scheduler/internal/api"
	"scheduler/internal/model"
	service "scheduler/internal/service_test"
	"testing"
)

func TestAPI_CancelJob(t *testing.T) {
	app := fiber.New()
	mockSvc := &service.MockJobService{}
	api.Register(app, mockSvc)

	t.Run("Valid cancellation", func(t *testing.T) {
		mockSvc.ExpectedCalls = nil
		jobID := "1"
		mockSvc.On("Cancel", jobID).Return(nil)

		req, err := http.NewRequest(fiber.MethodDelete, "/jobs/1", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Job not found", func(t *testing.T) {
		mockSvc.ExpectedCalls = nil
		jobID := "2"
		mockSvc.On("Cancel", jobID).Return(model.ErrJobNotFound)

		req, err := http.NewRequest(fiber.MethodDelete, "/jobs/2", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	t.Run("Cancel error", func(t *testing.T) {
		mockSvc.ExpectedCalls = nil
		jobID := "3"
		mockSvc.On("Cancel", jobID).Return(model.ErrJobAlreadyCancelled)

		req, err := http.NewRequest(fiber.MethodDelete, "/jobs/3", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
