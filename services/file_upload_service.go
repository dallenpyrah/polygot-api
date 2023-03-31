package services

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"polygot-api/contracts"
	"polygot-api/interfaces"
	"time"
)

type FileUploadService struct {
	fileLocationDetailsWriter interfaces.FileLocationDetailsWriter
}

func NewFileUploadService(fileLocationDetailsWriter interfaces.FileLocationDetailsWriter) FileUploadService {
	return FileUploadService{fileLocationDetailsWriter: fileLocationDetailsWriter}
}

func (f *FileUploadService) UploadFile(file *multipart.FileHeader) (contracts.UploadFileResponseContract, error) {
	var uploadFileResponseContract contracts.UploadFileResponseContract

	uniqueFileName := generateUniqueFileName(file.Filename)

	err := saveFile(file, "./uploads/"+uniqueFileName)
	if err != nil {
		return uploadFileResponseContract, err
	}

	requestId, err := f.fileLocationDetailsWriter.InsertFileLocationDetails(uniqueFileName)
	if err != nil {
		return uploadFileResponseContract, err
	}

	err = f.fileLocationDetailsWriter.InsertFileUploadStatus(requestId, "pending")
	if err != nil {
		return uploadFileResponseContract, err
	}

	uploadFileResponseContract.RequestId = requestId
	uploadFileResponseContract.FileName = uniqueFileName

	return uploadFileResponseContract, nil
}

func (f *FileUploadService) GetFileTranslationStatus(id int64) (contracts.FileUploadStatusContract, error) {
	var fileUploadStatusContract contracts.FileUploadStatusContract

	status, err := f.fileLocationDetailsWriter.GetFileUploadStatus(id)
	if err != nil {
		return fileUploadStatusContract, err
	}

	fileUploadStatusContract.Status = status

	return fileUploadStatusContract, nil
}

func generateUniqueFileName(fileName string) string {
	extension := filepath.Ext(fileName)

	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(1000)

	newFilename := fmt.Sprintf("%d-%d%s", timestamp, randomNumber, extension)

	return newFilename
}

func saveFile(file *multipart.FileHeader, destinationFile string) error {
	sourceFile, err := file.Open()
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create(destinationFile)

	if err != nil {
		return err
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)

	if err != nil {
		return err
	}

	return nil
}
