package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)
const LogsDocuments = "logsdocuments"
const ContentImage = "images"
const ContentDocuments = "application/pdf"

var theSession *session.Session

func GetSession() *session.Session {

	if theSession == nil {
		theSession = initSession()
	}

	return theSession
}

func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}))
	return newSession
}