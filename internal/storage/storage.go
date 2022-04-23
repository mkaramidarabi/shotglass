package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"
	"os"
	"path/filepath"
	"shotglass/internal/conf"
)

func PutObject(fileName string) (minio.UploadInfo, error) {
	var uploadInfo minio.UploadInfo
	minioClient, err := minio.New(
		conf.AppConfig.AwsEndpoint, &minio.Options{
			Creds: credentials.NewStaticV4(
				conf.AppConfig.AwsId,
				conf.AppConfig.AwsSecret,
				""),
			Secure: true,
		})
	if err != nil {
		fmt.Println(err)
		return uploadInfo, err
	}

	contentType, err := getFileContentType(fileName)
	if err != nil {
		fmt.Println(err)
		return uploadInfo, err
	}
	//https://docs.min.io/docs/golang-client-api-reference.html#FPutObject
	uploadInfo, err = minioClient.FPutObject(context.Background(),
		conf.AppConfig.AwsBucket,
		filepath.Base(fileName),
		fileName,
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	if err != nil {
		fmt.Println(err)
		return uploadInfo, err
	}
	return uploadInfo, nil
}

func getFileContentType(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
