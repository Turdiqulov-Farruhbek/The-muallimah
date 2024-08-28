package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gitlab.com/muallimah/course_service/internal/pkg/config"
)

type MinIO struct {
	client *minio.Client
	Cf *config.Config
}
var ContentType = map[string]string{
	".png": "image/png",
	".pdf": "application/pdf",
}

func MinIOConnect(cf *config.Config) (*MinIO, error) {
	endpoint := cf.MinioUrl
	accessKeyID := cf.MinioUser
	secretAccessKey := cf.MinioPassword
	useSSL := false // Use false if not using SSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Make a bucket if it does not exist
	bucketName := cf.BucketName

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil && exists {
			log.Println(err)
			return nil, err
		}
	}

	policy := fmt.Sprintf(`{
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Principal": "*",
              "Action": ["s3:GetObject"],
              "Resource": ["arn:aws:s3:::%s/*"]
          }
      ]
  }`, bucketName)

	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		log.Println("error while setting bucket policy : ", err)
		return nil, err
	}

	return &MinIO{
		client: minioClient,
		Cf: cf,
		}, err
}
func (m *MinIO) Upload(fileName string,contentType string) (*string,error ){

	uploadPath := m.Cf.MinioPath + fileName
	c_type := ContentType[contentType]
	_, err := m.client.FPutObject(context.Background(), m.Cf.BucketName, fileName, uploadPath, minio.PutObjectOptions{
		ContentType: c_type,
	})

	if err != nil {
		return nil,fmt.Errorf("error while uploading to minio: %v", err)
	}

	minioURL := fmt.Sprintf("http://localhost:9000/%s/%s",m.Cf.BucketName ,fileName)
	return &minioURL, nil
}
