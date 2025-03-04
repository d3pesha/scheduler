package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"scheduler/internal/api"
	"scheduler/internal/model"
	"scheduler/internal/service_test"
	"testing"
	"time"
)

func TestAPI_CreateJob(t *testing.T) {
	app := fiber.New()
	mockSvc := &service.MockJobService{}
	api.Register(app, mockSvc)

	t.Run("Valid request", func(t *testing.T) {
		mockJob := &model.Job{
			ID:          "test-job-id",
			Description: "Test Job",
			ExecuteAt:   time.Date(2044, 1, 1, 12, 0, 0, 0, time.UTC),
			Status:      model.StatusScheduled,
		}

		mockSvc.On("Create", mock.Anything, "Test Job",
			time.Date(2044, 1, 1, 12, 0, 0, 0, time.UTC)).
			Return(mockJob, nil)

		payload := `{"description": "Test Job", "executeAt": "2044-01-01T12:00:00Z"}`
		req, err := http.NewRequest(fiber.MethodPost, "/jobs", bytes.NewReader([]byte(payload)))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var job model.Job
		err = json.NewDecoder(resp.Body).Decode(&job)
		assert.NoError(t, err)
		assert.Equal(t, mockJob.ID, job.ID)
		assert.Equal(t, mockJob.Description, job.Description)
		assert.Equal(t, mockJob.ExecuteAt, job.ExecuteAt)
		assert.Equal(t, mockJob.Status, job.Status)
	})

	t.Run("Invalid request - missing description", func(t *testing.T) {
		mockSvc.On("Create", mock.Anything, "",
			time.Date(2044, 1, 1, 12, 0, 0, 0, time.UTC)).
			Return(&model.Job{}, model.ErrEmptyDescription)

		payload := `{"executeAt": "2044-01-01T12:00:00Z"}`
		req, err := http.NewRequest(fiber.MethodPost, "/jobs", bytes.NewReader([]byte(payload)))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Invalid request - invalid executeAt format", func(t *testing.T) {
		payload := `{"description": "Test Job", "executeAt": "invalid-date-format"}`
		req, err := http.NewRequest(fiber.MethodPost, "/jobs", bytes.NewReader([]byte(payload)))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
