package controllers

import "github.com/gofiber/fiber/v2"

type FolderController struct{}

func (f FolderController) UploadFolderForTranslation(request *fiber.Ctx) error {
	return request.SendString("UploadProjectForTranslation")
}

func (f FolderController) GetFolderTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}

func (f FolderController) GetFolderTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}
