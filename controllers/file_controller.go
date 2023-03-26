package controllers

import "github.com/gofiber/fiber/v2"

type FileController struct{}

func (fileController FileController) UploadFileForTranslation(request *fiber.Ctx) error {
	// Grab the file from the request
	file, err := request.FormFile("file")

	if err != nil {
		return request.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	err = request.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return request.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	return request.JSON(fiber.Map{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
	})
}

func (fileController FileController) GetFileTranslationResult(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}

func (fileController FileController) GetFileTranslationStatus(request *fiber.Ctx) error {
	return request.SendString("GetFileTranslation")
}
