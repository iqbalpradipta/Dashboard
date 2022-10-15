package helper

import (
	"fmt"
	"immersiveProject/config"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(directory string, fileName string, contentType string, fileData multipart.File) (string, error) {
	session := config.GetSession()
	uploader := s3manager.NewUploader(session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: 		aws.String(os.Getenv("AWS_BUCKET")),
		Key: 			aws.String("/" + directory + "/" + fileName),
		Body: 			fileData,
		ContentType: 	aws.String(contentType),
	})

	if err != nil{
		return "", fmt.Errorf("failed to upload file")
	}

	return result.Location, nil
}

func CheckfileExtension(filename string, contentType string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if contentType == config.ContentImage {
		if extension != "jpg" && extension != "jpeg" && extension != "png"{
			return "", fmt.Errorf("forbidden type data")
		}
	}

	return extension, nil
}

func CheckFileSize(size int64, contentType string) error {
	if size == 0 {
		return fmt.Errorf("ilegal file size or size 0 data")
	}
	if contentType == config.ContentImage{
		if size > 1097152 {
			return fmt.Errorf("file size maks 15 mb")
		}
	}
	if contentType == config.ContentDocuments{
		if size > 10097152 {
			return fmt.Errorf("file size maks 15 mb")
		}
	}
	return nil
}

func UploadPDFToS3(directory string, filename string, contentType string, data io.Reader) (string, error) {
	session := config.GetSession()
	uploader := s3manager.NewUploader(session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:  		aws.String(os.Getenv("AWS_BUCKET")),
		Key: 	 		aws.String("/" + directory + "/" + filename),
		Body: 	 		data,
		ContentType: 	aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file")
	}
	return result.Location, nil
}