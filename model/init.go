package model

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var gormDb *gorm.DB
var minioClient *minio.Client
var bucketName string

func Connect(dialector gorm.Dialector) {
	var err error
	gormDb, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if gormDb == nil {
		panic("DB is nil")
	}
}

func CreateTables() {
	if gormDb == nil {
		panic("DB is nil")
	}
	err := gormDb.AutoMigrate(
		&User{},
		&Organization{},
		&Department{},
		&JoinedDepartment{},
		&Event{},
		&Interview{},
		&JoinedInterview{},
		&Message{},
		&Image{})
	if err != nil {
		panic(err)
	}
}

func CreateRow(request interface{}) error {
	return gormDb.Create(request).Error
}

func ConnectObjectStorage() {
	bucketName = viper.GetString("minio.bucket_name")

	var err error
	minioClient, err = minio.New(viper.GetString("minio.endpoint"), &minio.Options{
		Creds: credentials.NewStaticV4(viper.GetString("minio.id"), viper.GetString("minio.secret"), ""),
		Secure: viper.GetBool("minio.secure"),
	})
	if err != nil {
		panic(err)
	}
	_, err = minioClient.ListBuckets(context.Background())
	if err != nil {
		panic(err)
	}

	err = createBuckets(bucketName)
	if err != nil {
		panic(err)
	}
}

func createBuckets(name string) error {
	if ok, err := minioClient.BucketExists(context.Background(), name); ok {
		return nil
	} else if err != nil {
		return err
	}

	err := minioClient.MakeBucket(context.Background(), name, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	return nil
}
