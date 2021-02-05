package database

import (
	"context"
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm/clause"
	"io"
	"mime/multipart"
)

type Image struct {
	ID           uint `gorm:"autoIncrement;primaryKey"`
	UserID       uint
	User         model.User
	OriginalName string
	CurrentName  string
}

func (image *Image) Save() error {
	return DB.Save(image).Error
}

func (image *Image) GetByUid() error {
	return DB.Preload(clause.Associations).Where(&model.User{ID: image.UserID}).First(image).Error
}

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
