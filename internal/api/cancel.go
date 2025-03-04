package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"scheduler/internal/model"
)

func (h *handler) cancelJob(c *fiber.Ctx) error {
	jobID := c.Params("id")

	err := h.jobService.Cancel(jobID)
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

	return c.SendStatus(fiber.StatusOK)
}
