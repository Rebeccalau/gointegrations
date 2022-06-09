package awsintegration

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Upload() {
	MaxRetries := 10
	UserAcc := true
	path, _ := os.Getwd()
	filename := filepath.Join(path, "data", "test_upload.txt")
	Region := "eu-west-1"

	bucket := "bucket-name"
	SecretAccessKey := "secret-access-key"
	AccessKeyID := "access-id"
	Token := "token"
	Key := "key"

	httpclient := getHttpClient(false)

	sess, sessionError := session.NewSession(
		&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				Token,
			),
			MaxRetries:      &MaxRetries,
			S3UseAccelerate: &UserAcc,
			HTTPClient:      httpclient,
			Region:          aws.String(Region),
		})

	if sessionError != nil {
		fmt.Println(sessionError)
	}

	uploader := s3manager.NewUploader(sess)

	f, fileErr := os.Open(filename)

	if fileErr != nil {
		fmt.Println(fileErr)
	}

	_, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(Key),
		Body:   f,
	})

	if uploadError != nil {
		fmt.Println(uploadError)
	}
}

func getHttpClient(proxy bool) *http.Client {
	if proxy {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: func(*http.Request) (*url.URL, error) {
					return url.Parse("http://user:pass@proxy:8080")
				},
			},
		}
	}

	return &http.Client{Transport: &http.Transport{}}
}
