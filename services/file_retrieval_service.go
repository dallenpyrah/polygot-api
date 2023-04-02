package services

import (
	"polygot-api/contracts"
	"polygot-api/interfaces"
)

type FileRetrievalService struct {
	fileRetrievalReader interfaces.FileRetrievalReader
}

func NewFileRetrievalService(fileRetrievalReader interfaces.FileRetrievalReader) FileRetrievalService {
	return FileRetrievalService{fileRetrievalReader: fileRetrievalReader}
}

func (f *FileRetrievalService) GetFileTranslationStatus(id int64) (contracts.FileUploadStatusContract, error) {
	var fileUploadStatusContract contracts.FileUploadStatusContract

	status, err := f.fileRetrievalReader.GetFileUploadStatus(id)
	if err != nil {
		return fileUploadStatusContract, err
	}

	fileUploadStatusContract.Status = status

	return fileUploadStatusContract, nil
}

func (f *FileRetrievalService) GetFileTranslationResult(id int64) (contracts.FileTranslationResultContract, error) {
	var fileTranslationResultContract contracts.FileTranslationResultContract

	fileTranslationResultContract.Id = id
	fileTranslationResultContract.FileName = "translated-"
	fileTranslationResultContract.Content = "Translated content"

	return fileTranslationResultContract, nil
}
