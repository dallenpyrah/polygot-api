package controllers

import "github.com/gofiber/fiber/v2"

type ProjectController struct{}

func (projectController ProjectController) UploadProjectForTranslation(request *fiber.Ctx) error {
	return request.SendString("UploadProjectForTranslation")
}

func (projectController ProjectController) GetProjectTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}

func (projectController ProjectController) GetProjectTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}
