package main

import (
	"centralize-jobs/crawlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/github", func(c *fiber.Ctx) error {
		return c.JSON(crawlers.GetJobsGithubBackendBr())
	})

	app.Get("/programathor", func(c *fiber.Ctx) error {
		return c.JSON(crawlers.GetJobsProgramathor())
	})

	app.Get("/vulpi", func(c *fiber.Ctx) error {
		return c.JSON(crawlers.GetJobsVulpi())
	})

	app.Listen(":3000")
}
