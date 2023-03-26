package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"polygot-api/controllers"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("polygot-api is running")
		return err
	})

	app.Get("/api/v1/languages", controllers.GetAllSupportedLanguages)
	app.Get("/api/v1/project/:id", controllers.GetProjectTranslationResult)
	app.Get("/api/v1/file/:id", controllers.GetFileTranslationResult)
	app.Get("/api/v1/folder/:id", controllers.GetFolderTranslationResult)
	app.Get("/api/v1/project/:id/status", controllers.GetProjectTranslationStatus)
	app.Get("/api/v1/file/:id/status", controllers.GetFileTranslationStatus)
	app.Get("/api/v1/folder/:id/status", controllers.GetFolderTranslationStatus)
	app.Post("/api/v1/file", controllers.UploadFileForTranslation)
	app.Post("/api/v1/project", controllers.UploadProjectForTranslation)
	app.Post("/api/v1/folder", controllers.UploadFolderForTranslation)

	app.Listen(":3000")
}
