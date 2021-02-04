package database

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
)

func CreateFile(ctx context.Context, objectName, contentType string, file multipart.File, objectSize int64) error {
	/* set file pointer to the start of the file */
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	/* store the file into MinIO storage */
	_, err = minioClient.PutObject(ctx, bucketName, objectName, file, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetFile(ctx context.Context, filePath string) (*minio.Object, error) {
	object, err := minioClient.GetObject(ctx, bucketName, filePath, minio.GetObjectOptions{})
	return object, err
}
