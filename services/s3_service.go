package services

import (
	"mime/multipart"
	"polygot-api/providers"
)

type S3Service struct {
	s3Provider providers.S3Provider
}

func (s S3Service) UploadFileToS3(fileHeader *multipart.FileHeader, fileName string) error {
	file, err := fileHeader.Open()

	bucketName := "bucket-name"

	if err != nil {
		return err
	}

	s.s3Provider.UploadFileToS3Bucket(bucketName, fileName, file)

	return nil
}
