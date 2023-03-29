package controllers

import "github.com/gofiber/fiber/v2"

type ProjectController struct{}

func (p ProjectController) UploadProjectForTranslation(request *fiber.Ctx) error {
	return request.SendString("UploadProjectForTranslation")
}

func (p ProjectController) GetProjectTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}

func (p ProjectController) GetProjectTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetProjectTranslation")
}
