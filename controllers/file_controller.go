package controllers

import "github.com/gofiber/fiber/v2"

func UploadFileForTranslation(c *fiber.Ctx) error {
	// Grab the file from the request
	file, err := c.FormFile("file")
	if err != nil {

	}

	// Save the file to the server
	err = c.SaveFile(file, "./uploads/"+file.Filename)

	return c.SendString("UploadFileForTranslation")
}

func GetFileTranslationResult(c *fiber.Ctx) error {
	return c.SendString("GetFileTranslation")
}

func GetFileTranslationStatus(c *fiber.Ctx) error {
	return c.SendString("GetFileTranslation")
}
