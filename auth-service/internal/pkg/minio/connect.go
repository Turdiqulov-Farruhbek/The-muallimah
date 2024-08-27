package minio_connect

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client     *minio.Client
	bucketName string
}

func NewMinioClient() (*MinioClient, error) {
	endpoint := "localhost:9000"
	accessKeyID := "user"
	secretAccessKey := "password"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println("Error creating Minio client:", err)
		return nil, err
	}

	bucketName := "pfp-bucket"

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil {
			return nil, errBucketExists
		}
		if exists {
			log.Println("Bucket already exists:", bucketName)
		} else {
			log.Println("Error creating bucket:", err)
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
		log.Println("Error while setting bucket policy:", err)
		return nil, err
	}

	return &MinioClient{Client: minioClient, bucketName: bucketName}, nil
}

func (m *MinioClient) DefaultBucket() string {
	return m.bucketName
}
