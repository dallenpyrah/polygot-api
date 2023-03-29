package controllers

import (
	"github.com/gofiber/fiber/v2"
	"polygot-api/services"
)

type FileController struct {
	fileUploadService services.FileUploadService
}

func (f FileController) UploadFileForTranslation(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to process the uploaded file",
		})
	}

	result, err := f.fileUploadService.UploadFile(file)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "An error occurred while uploading the file",
		})
	}

	return c.Status(200).JSON(result)
}

func (f FileController) GetFileTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}

func (f FileController) GetFileTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}
