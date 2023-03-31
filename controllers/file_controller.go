package controllers

import (
	"github.com/gofiber/fiber/v2"
	"polygot-api/services"
	"strconv"
)

type FileController struct {
	fileUploadService services.FileUploadService
}

func NewFileController(fileUploadService services.FileUploadService) FileController {
	return FileController{fileUploadService: fileUploadService}
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
	id := request.Params("id")
	fileUploadId, _ := strconv.ParseInt(id, 10, 64)

	result, err := f.fileUploadService.GetFileTranslationStatus(fileUploadId)
	if err != nil {
		return request.Status(500).JSON(fiber.Map{
			"error": "An error occurred while getting the file translation status",
		})
	}

	return request.Status(200).JSON(result)
}
