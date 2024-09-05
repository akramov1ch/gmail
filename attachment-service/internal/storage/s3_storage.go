package storage

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Session *session.Session
var s3Service *s3.S3

func InitS3() {
	s3Session = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	}))
	s3Service = s3.New(s3Session)
}

func UploadFile(file os.File, filename string, contentType string) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	file.Seek(0, 0)

	_, err = s3Service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(filename),
		Body:        &file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET"), filename)
	return fileURL, nil
}

func DeleteFile(filename string) error {
	_, err := s3Service.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(filename),
	})
	if err != nil {
		return err
	}
	return nil
}
