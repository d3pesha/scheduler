package api

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h *handler) createJob(c *fiber.Ctx) error {
	var req struct {
		Description string    `json:"description"`
		ExecuteAt   time.Time `json:"executeAt"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	job, err := h.jobService.Create(c.Context(), req.Description, req.ExecuteAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(job)
}
