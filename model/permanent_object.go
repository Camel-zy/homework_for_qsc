package model

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	_ "gorm.io/gorm/clause"
)

type PermanentObject struct {
	ID         uint      `gorm:"autoIncrement;primaryKey"`
	UUID       uuid.UUID `gorm:"not null;type:uuid;uniqueIndex"`
	ObjectName string    `gorm:"not null"`
	Finished   bool      `gorm:"not null;default:false"`
}

func CreateObject(ctx context.Context, objectName string) (url *url.URL, formData map[string]string, id uuid.UUID, err error) {
	policy := minio.NewPostPolicy()
	uuid := uuid.NewV4()
	_ = policy.SetBucket(bucketName)
	_ = policy.SetKey(uuid.String())
	_ = policy.SetExpires(time.Now().UTC().Add(time.Minute * 15))
	_ = policy.SetContentLengthRange(0, 10*1024*1024) // up to 10 MiB
	url, formData, err = minioClient.PresignedPostPolicy(ctx, policy)
	if err != nil {
		logrus.Error(err)
		return
	}
	permanentObject := PermanentObject{
		UUID:       uuid,
		ObjectName: objectName,
		Finished:   false,
	}
	result := gormDb.Create(&permanentObject)
	return url, formData, uuid, result.Error
}

func SealObject(ctx context.Context, uuid uuid.UUID) error {
	result := gormDb.Model(&PermanentObject{}).Where("UUID = ?", uuid.String()).Updates(&PermanentObject{Finished: true})
	return result.Error
}

func GetObject(ctx context.Context, uuid uuid.UUID) (presignedURL *url.URL, err error) {
	var permanentObject PermanentObject
	result := gormDb.Model(&PermanentObject{UUID: uuid}).Where("UUID = ?", uuid.String()).First(&permanentObject)
	if result.Error != nil {
		logrus.Error(err)
		return nil, err
	}
	if !permanentObject.Finished {
		err := fmt.Errorf("Object with uuid %s does not exist or is not ready", uuid)
		logrus.Error(err)
		return nil, err
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf(`attachment; filename="%s"`, permanentObject.ObjectName))

	presignedURL, err = minioClient.PresignedGetObject(ctx,
		bucketName, permanentObject.UUID.String(), time.Minute*10, reqParams)
	if err != nil {
		fmt.Println(err)
	}
	return
}
