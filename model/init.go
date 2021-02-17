package model

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Brief struct {
	ID              uint
	Name            string
	Description     string
}

var ErrNoRowsAffected = errors.New("no rows affected")
var ErrInternalError = errors.New("internal error")

var gormDb *gorm.DB
var minioClient *minio.Client
var bucketName string

func Connect(dialector gorm.Dialector) {
	var err error
	gormDb, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	if gormDb == nil {
		logrus.Fatal("DB is nil")
	}

	logrus.Info("PostgreSQL connected")
}

func CreateTables() {
	if gormDb == nil {
		logrus.Fatal("DB is nil")
	}
	err := gormDb.AutoMigrate(
		&User{},
		&Organization{},
		&OrganizationHasUser{},
		&Department{},
		&JoinedDepartment{},
		&Event{},
		&Interview{},
		&JoinedInterview{},
		&Message{},
		&MessageTemplate{},
		&Image{})
	if err != nil {
		logrus.Fatal(err)
	}
}

func CreateRow(request interface{}) error {
	return gormDb.Create(request).Error
}

func ConnectObjectStorage() {
	if !viper.GetBool("minio.enable") {
		return
	}

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
		logrus.Fatal(err)
	}

	err = createBuckets(bucketName)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("MinIO connected")
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
