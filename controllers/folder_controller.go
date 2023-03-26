package controllers

import "github.com/gofiber/fiber/v2"

func UploadFolderForTranslation(c *fiber.Ctx) error {
	return c.SendString("UploadProjectForTranslation")
}

func GetFolderTranslationResult(c *fiber.Ctx) error {
	return c.SendString("GetProjectTranslation")
}

func GetFolderTranslationStatus(c *fiber.Ctx) error {
	return c.SendString("GetProjectTranslation")
}
