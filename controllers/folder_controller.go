package controllers

import "github.com/gofiber/fiber/v2"

type FolderController struct{}

func (folderController FolderController) UploadFolderForTranslation(request *fiber.Ctx) error {
	return request.SendString("UploadProjectForTranslation")
}

func (folderController FolderController) GetFolderTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}

func (folderController FolderController) GetFolderTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}
