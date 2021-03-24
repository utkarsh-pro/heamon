package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utkarsh-pro/heamon/models"
	"github.com/utkarsh-pro/heamon/pkg/utils"
)

// NewRoutes registers the REST endpoints to the fiber app
func NewRoutes(app *fiber.App, handlers models.Handler) {
	app.Put("/api/v1/config", handlers.RegisterNewConfig)
	app.Get("/api/v1/config", handlers.GetConfig)
	app.Get("/api/v1/status", handlers.GetStatus)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": utils.GetEnv("APP_NAME", "Heamon"),
		})
	})
	app.Static("/", "./ui/build")
}