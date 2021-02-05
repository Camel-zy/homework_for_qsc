package database

import (
	"context"
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB
var minioClient *minio.Client
var bucketName string

func Connect(dialector gorm.Dialector) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("DB is nil")
	}
	DB = db
}

func CreateTables() {
	if DB == nil {
		panic("DB is nil")
	}
	err := DB.AutoMigrate(
		&model.User{},
		&model.Organization{},
		&model.Department{},
		&model.JoinedDepartment{},
		&model.Event{},
		&model.Interview{},
		&model.JoinedInterview{},
		&model.Message{},
		&Image{})
	if err != nil {
		panic(err)
	}
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
