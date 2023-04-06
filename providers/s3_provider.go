package providers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"mime/multipart"
)

type S3Provider struct {
}

func (s S3Provider) CreateNewS3Client() *s3.S3 {
	accessKey := "YOUR_AWS_ACCESS_KEY"
	secretKey := "YOUR_AWS_SECRET_KEY"
	region := "YOUR_AWS_REGION"

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	serviceClient := s3.New(session)

	return serviceClient
}

func (s S3Provider) UploadFileToS3Bucket(bucketName string, fileName string, file multipart.File) error {
	serviceClient := s.CreateNewS3Client()

	_, err := serviceClient.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   aws.ReadSeekCloser(file),
	})

	if err != nil {
		return err
	}

	return nil
}
