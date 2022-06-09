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
	//filename := filepath.Join(path, "data", "video.mp4")
	filename := filepath.Join(path, "data", "test_upload.txt")
	Region := "eu-west-1"

	Bucket := os.Getenv("AWS_Bucket")
	SecretAccessKey := os.Getenv("AWS_SecretAccessKey")
	AccessKeyID := os.Getenv("AWS_AccessKeyID")
	Token := os.Getenv("AWS_Token")
	Key := os.Getenv("AWS_Key")

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
		return
	}

	uploader := s3manager.NewUploader(sess)

	f, fileErr := os.Open(filename)

	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	fileInfo, fileStatError := f.Stat()
	if fileStatError != nil {
		fmt.Println(fileStatError)
		return
	}

	reader := &CustomReader{
		fp:      f,
		size:    fileInfo.Size(),
		signMap: map[int64]struct{}{},
	}

	_, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(Key),
		Body:   reader,
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
