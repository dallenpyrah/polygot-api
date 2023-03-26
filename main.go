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

	app.Post("/api/v1/project", controllers.ProjectController{}.UploadProjectForTranslation)
	app.Get("/api/v1/project/:id/status", controllers.ProjectController{}.GetProjectTranslationStatus)
	app.Get("/api/v1/project/:id", controllers.ProjectController{}.GetProjectTranslationResult)

	app.Get("/api/v1/file/:id", controllers.FileController{}.GetFileTranslationResult)
	app.Get("/api/v1/file/:id/status", controllers.FileController{}.GetFileTranslationStatus)
	app.Post("/api/v1/file", controllers.FileController{}.UploadFileForTranslation)

	app.Post("/api/v1/folder", controllers.FolderController{}.UploadFolderForTranslation)
	app.Get("/api/v1/folder/:id", controllers.FolderController{}.GetFolderTranslationResult)
	app.Get("/api/v1/folder/:id/status", controllers.FolderController{}.GetFolderTranslationStatus)

	app.Listen(":3000")
}
