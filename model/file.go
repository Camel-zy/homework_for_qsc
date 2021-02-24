package model

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"net/url"
	"time"
)

type Image struct {
	ID           uint `gorm:"autoIncrement;primaryKey"`
	UserID       uint
	User         User
	OriginalName string
	CurrentName  string
}

func (image *Image) Save() error {
	return gormDb.Save(image).Error
}

func (image *Image) GetByUid() error {
	return gormDb.Preload(clause.Associations).Where(&User{ID: image.UserID}).First(image).Error
}

func CreateFile(ctx context.Context, objectName string,
	userMetadata map[string]string) (url *url.URL, formData map[string]string, err error) {
	policy := minio.NewPostPolicy()
	for k, v := range userMetadata {
		_ = policy.SetUserMetadata(k, v)
	}
	_ = policy.SetBucket(bucketName)
	_ = policy.SetKey(objectName)
	_ = policy.SetExpires(time.Now().UTC().Add(time.Minute * 5))
	// _ = policy.SetContentType("image/*")
	_ = policy.SetContentLengthRange(1024, 1024 * 1024)  // from 1 KB to 1 MB
	url, formData, err = minioClient.PresignedPostPolicy(ctx, policy)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func GetFile(ctx context.Context, objectName string) (presignedURL *url.URL, err error) {
	uuidString := uuid.NewV4().String()
	reqParams := make(url.Values)
	// reqParams.Set("response-content-type", "image/jpeg")
	reqParams.Set("response-content-disposition", fmt.Sprintf(`attachment; filename="%s"`,uuidString))

	presignedURL, err = minioClient.PresignedGetObject(ctx,
		bucketName, objectName, time.Minute * 10, reqParams)
	if err != nil {
		fmt.Println(err)
	}
	return
}
