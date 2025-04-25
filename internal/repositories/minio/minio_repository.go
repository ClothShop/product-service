package minio

import (
	"context"
	"fmt"
	minio2 "github.com/ClothShop/product-service/internal/config/minio"
	"github.com/ClothShop/product-service/internal/dto"
	minio "github.com/minio/minio-go/v7"
	"io"
	"os"
	"time"
)

func UploadFile(product *dto.ProductCreate) []string {
	if product.Images == nil || len(product.Images) == 0 {
		return nil
	}

	var uploadedImages []string

	for _, image := range product.Images {
		imageFile, err := image.Open()
		if err != nil {
			return nil
		}

		objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), image.Filename)
		bucket := os.Getenv("MINIO_BUCKET")
		contentType := image.Header.Get("Content-Type")

		_, err = minio2.Client.PutObject(
			context.Background(),
			bucket,
			objectName,
			imageFile,
			image.Size,
			minio.PutObjectOptions{ContentType: contentType},
		)
		imageFile.Close()

		if err != nil {
			return nil
		}
		uploadedImages = append(uploadedImages, objectName)
	}

	return uploadedImages
}

func GetFile(objectName string) ([]byte, error) {
	bucket := os.Getenv("MINIO_BUCKET")

	object, err := minio2.Client.GetObject(
		context.Background(),
		bucket,
		objectName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}
	defer object.Close()

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DeleteFile(objectName string) error {
	bucket := os.Getenv("MINIO_BUCKET")

	err := minio2.Client.RemoveObject(
		context.Background(),
		bucket,
		objectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetFiles(objectNames []string) ([][]byte, error) {
	var files [][]byte
	for _, name := range objectNames {
		file, err := GetFile(name)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func DeleteFiles(objectNames []string) error {
	for _, name := range objectNames {
		err := DeleteFile(name)
		if err != nil {
			return err
		}
	}
	return nil
}
