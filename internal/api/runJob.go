package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"scheduler/internal/model"
)

func (h *handler) runJob(c *fiber.Ctx) error {
	jobID := c.Params("id")

	job, err := h.jobService.RunJob(c.Context(), jobID)
	if err != nil {
		if errors.Is(err, model.ErrJobNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(job)
}
