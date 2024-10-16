package vender

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	StorageClient           = *minio.Client
	StorageUploadInfo       = minio.UploadInfo
	StoragePutObjectOptions = minio.PutObjectOptions
	StorageGetObjectOptions = minio.GetObjectOptions
)

type StorageConfig struct {
	EndPoint        string
	AccessKey       string
	SecretAccessKey string
}

func NewObjectStorage(config *StorageConfig) (StorageClient, error) {
	const (
		useSSL = false
	)
	var (
		endpoint        = config.EndPoint
		accessKeyID     = config.AccessKey
		secretAccessKey = config.SecretAccessKey
	)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
