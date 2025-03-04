package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"scheduler/internal/model"
)

func (h *handler) getAllJobs(c *fiber.Ctx) error {
	jobs, err := h.jobService.GetAll()
	if err != nil {
		if errors.Is(err, model.ErrJobNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(jobs)
}
