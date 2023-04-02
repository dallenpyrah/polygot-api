package controllers

import (
	"github.com/gofiber/fiber/v2"
	"polygot-api/services"
	"strconv"
)

type FileController struct {
	fileUploadService    services.FileUploadService
	fileRetrievalService services.FileRetrievalService
}

const (
	StatusBadRequest          = 400
	StatusInternalServerError = 500
	StatusOK                  = 200
)

func NewFileController(fileUploadService services.FileUploadService, fileRetrievalService services.FileRetrievalService) FileController {
	return FileController{fileUploadService: fileUploadService, fileRetrievalService: fileRetrievalService}
}

func (f FileController) UploadFileForTranslation(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to process the uploaded file",
		})
	}

	result, err := f.fileUploadService.UploadFile(file)

	if err != nil {
		return c.Status(StatusInternalServerError).JSON(fiber.Map{
			"error": "An error occurred while uploading the file",
		})
	}

	return c.Status(StatusOK).JSON(result)
}

func (f FileController) GetFileTranslationResult(request *fiber.Ctx) error {
	id := request.Params("id")
	fileUploadId, _ := strconv.ParseInt(id, 10, 64)

	result, err := f.fileRetrievalService.GetFileTranslationResult(fileUploadId)
	if err != nil {
		return request.Status(StatusInternalServerError).JSON(fiber.Map{
			"error": "An error occurred while getting the file translation result",
		})
	}

	return request.Status(StatusOK).JSON(result)
}

func (f FileController) GetFileTranslationStatus(request *fiber.Ctx) error {
	id := request.Params("id")
	fileUploadId, _ := strconv.ParseInt(id, 10, 64)

	result, err := f.fileRetrievalService.GetFileTranslationStatus(fileUploadId)
	if err != nil {
		return request.Status(StatusInternalServerError).JSON(fiber.Map{
			"error": "An error occurred while getting the file translation status",
		})
	}

	return request.Status(StatusOK).JSON(result)
}
