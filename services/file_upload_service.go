package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"path/filepath"
	"polygot-api/interfaces"
	"time"
)

type FileUploadService struct {
	fileLocationDetailsWriter interfaces.FileLocationDetailsWriter
}

func (f *FileUploadService) UploadFile(request *fiber.Ctx) error {
	file, err := request.FormFile("file")

	if err != nil {
		return request.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	file.Filename = GenerateUniqueFileName(file.Filename)

	err = request.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return request.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	requestId := f.fileLocationDetailsWriter.InsertFileLocationDetails(file.Filename)

	return request.JSON(fiber.Map{
		"message":   "File uploaded successfully",
		"filename":  file.Filename,
		"requestId": requestId,
	})
}

func GenerateUniqueFileName(fileName string) string {
	extension := filepath.Ext(fileName)

	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(1000)

	newFilename := fmt.Sprintf("%d-%d%s", timestamp, randomNumber, extension)

	return newFilename
}
