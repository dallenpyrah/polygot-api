package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllSupportedLanguages(c *fiber.Ctx) error {
	return c.SendString("GetAllSupportedLanguages")
}
