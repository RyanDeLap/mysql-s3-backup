package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager" 
	"os"
	"fmt"
)

func Upload(fileName string) (string, error) {
	bucket := os.Getenv("BUCKET")
	uploadPath := os.Getenv("UPLOAD_PATH")
	region := os.Getenv("REGION")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{Region: aws.String(region)},
	}))

	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", fileName, err)
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(uploadPath + fileName),
		Body:   file,
	})


	return result.Location, err
}
