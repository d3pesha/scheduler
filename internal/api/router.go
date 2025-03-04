package api

import (
	"github.com/gofiber/fiber/v2"
	"scheduler/internal/service"
)

type handler struct {
	jobService service.Service
}

func Register(r fiber.Router, jobService service.Service) {
	h := &handler{
		jobService: jobService,
	}

	api := r.Group("/jobs")

	api.Post("/", h.createJob)
	api.Get("/", h.getAllJobs)
	api.Get("/:id", h.getJobByID)
	api.Delete("/:id", h.cancelJob)
	api.Post("/:id/run", h.runJob)
}
