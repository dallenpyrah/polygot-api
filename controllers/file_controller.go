package controllers

import (
	"github.com/gofiber/fiber/v2"
	"polygot-api/services"
)

type FileController struct {
	fileUploadService services.FileUploadService
}

func (f FileController) UploadFileForTranslation(request *fiber.Ctx) error {
	result := f.fileUploadService.UploadFile(request)
	return result
}

func (f FileController) GetFileTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}

func (f FileController) GetFileTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}
