package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"scheduler/internal/api"
	"scheduler/internal/model"
	service "scheduler/internal/service_test"
	"testing"
)

func TestAPI_GetAllJobs(t *testing.T) {
	app := fiber.New()
	mockSvc := &service.MockJobService{}
	api.Register(app, mockSvc)

	t.Run("Get all jobs - success", func(t *testing.T) {
		mockSvc.ExpectedCalls = nil
		mockSvc.On("GetAll").Return([]*model.Job{
			{ID: "1", Description: "Job 1"},
			{ID: "2", Description: "Job 2"},
		}, nil)

		req, err := http.NewRequest(fiber.MethodGet, "/jobs", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var jobs []model.Job
		err = json.NewDecoder(resp.Body).Decode(&jobs)
		assert.NoError(t, err)
		assert.Len(t, jobs, 2)
		assert.Equal(t, "Job 1", jobs[0].Description)
	})

	t.Run("Get all jobs - not found", func(t *testing.T) {
		mockSvc.ExpectedCalls = nil
		mockSvc.On("GetAll").Return([]*model.Job{}, model.ErrJobNotFound)

		req, err := http.NewRequest(fiber.MethodGet, "/jobs", nil)
		assert.NoError(t, err)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		var body map[string]string
		err = json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)

		assert.Equal(t, "job not found", body["error"])
	})
}
