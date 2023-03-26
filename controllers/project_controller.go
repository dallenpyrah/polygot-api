package controllers

import "github.com/gofiber/fiber/v2"

func UploadProjectForTranslation(c *fiber.Ctx) error {
	return c.SendString("UploadProjectForTranslation")
}

func GetProjectTranslationResult(c *fiber.Ctx) error {
	return c.SendString("GetProjectTranslation")
}

func GetProjectTranslationStatus(c *fiber.Ctx) error {
	return c.SendString("GetProjectTranslation")
}
