package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"polygot-api/controllers"
	"polygot-api/providers"
	"polygot-api/repositories"
	"polygot-api/services"
)

func main() {
	app := fiber.New()

	os.Setenv("FIREBASE_CREDENTIALS", "/secrets/firebase-storage-service-account.json")

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("polygot-api is running")
		return err
	})

	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		err := c.SendString("polygot-api is healthy")
		return err
	})

	app.Post("/api/v1/project", controllers.ProjectController{}.UploadProjectForTranslation)
	app.Get("/api/v1/project/:id/status", controllers.ProjectController{}.GetProjectTranslationStatus)
	app.Get("/api/v1/project/:id", controllers.ProjectController{}.GetProjectTranslationResult)

	fileLocationsDetailsRepository := repositories.NewFileLocationDetailsRepository(providers.DbConnectionProvider{}, nil)
	fileUploadService := services.NewFileUploadService(fileLocationsDetailsRepository)
	fileController := controllers.NewFileController(fileUploadService)

	app.Get("/api/v1/file/:id", fileController.GetFileTranslationResult)
	app.Get("/api/v1/file/:id/status", fileController.GetFileTranslationStatus)
	app.Post("/api/v1/file", fileController.UploadFileForTranslation)

	app.Post("/api/v1/folder", controllers.FolderController{}.UploadFolderForTranslation)
	app.Get("/api/v1/folder/:id", controllers.FolderController{}.GetFolderTranslationResult)
	app.Get("/api/v1/folder/:id/status", controllers.FolderController{}.GetFolderTranslationStatus)

	app.Listen(":3000")
}
