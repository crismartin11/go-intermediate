package test

import (
	"go-intermediate/pkg/storage"
	"testing"
)

func TestListBuckets(t *testing.T) {

	s3Client := storage.NewS3Client()
	_, err := s3Client.ListBuckets()

	if err != nil {
		t.Errorf("Error getting buckets: %v", err)
	}
}
