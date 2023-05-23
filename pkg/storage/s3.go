package storage

import (
	"fmt"
	"os"

	"go-intermediate/pkg/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct{}

func NewS3Client() S3Client {
	return S3Client{}
}

type Bucket struct {
	Name string `json:"Name"`
}

type BucketObj struct {
	Name string `json:"Name"`
}

func (s3c S3Client) ListBuckets() ([]Bucket, error) {
	service := credentials.GetClientS3()
	buckets := []Bucket{}

	// Usando la instancia del servicio, listo los buckets
	result, err := service.ListBuckets(nil)
	if err != nil {
		return buckets, fmt.Errorf("No se pudo listar buckets: %s", err)
	}

	for _, bucket := range result.Buckets {
		bucket := Bucket{Name: aws.StringValue(bucket.Name)}
		buckets = append(buckets, bucket)
	}

	return buckets, nil
}

func (s3c S3Client) ListObjects(bucketName string) ([]BucketObj, error) {
	service := credentials.GetClientS3()

	objects := []BucketObj{}
	result, err := service.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return objects, fmt.Errorf("No se pudo listar objetos: %s", err)
	}

	for _, obj := range result.Contents {
		object := BucketObj{Name: aws.StringValue(obj.Key)}
		objects = append(objects, object)
	}

	return objects, nil
}

func (s3c S3Client) CreateBucket(bucketName *string) (bool, error) {
	service := credentials.GetClientS3()

	_, err := service.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucketName,
	})
	if err != nil {
		return false, fmt.Errorf("No se pudo crear el bucket: %s", err)
	}

	err = service.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: bucketName,
	})
	if err != nil {
		return false, fmt.Errorf("No fue posible crear el bucket: %s", err)
	}

	return true, nil
}

func (s3c S3Client) Upload(bucketname *string, filename *string) (bool, error) {
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	service := s3manager.NewUploader(session)

	wd, err := os.Getwd()
	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	file, err := os.Open(wd + "/pkg/storage/filetest.json")
	if err != nil {
		return false, fmt.Errorf("No es posible leer el archivo: %s", err)
	}

	defer file.Close()

	_, err = service.Upload(&s3manager.UploadInput{
		Bucket: bucketname,
		Key:    filename,
		Body:   file,
	})
	if err != nil {
		return false, fmt.Errorf("No fue posible subir el archivo: %s", err)
	}

	return true, nil
}
