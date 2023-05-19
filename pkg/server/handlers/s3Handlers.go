package handlers

import (
	"net/http"

	"go-intermediate/pkg/storage"
)

type customDataS3 struct {
	Buckets  []storage.Bucket
	Objects  []storage.BucketObj
	Created  bool
	Uploaded bool
	Err      string
	commonData
}

var (
	dataS3 = customDataS3{
		Buckets:  []storage.Bucket{},
		Objects:  []storage.BucketObj{},
		Created:  false,
		Uploaded: false,
		Err:      "",
	}
)

func S3Handler(response http.ResponseWriter, request *http.Request) {
	dataS3.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "s3.gohtml", dataS3)
}

func S3BucketsHandler(response http.ResponseWriter, request *http.Request) {

	s3Client := storage.NewS3Client()
	listBuckets, err := s3Client.ListBuckets()

	if err != nil {
		dataS3.Err = err.Error()
	} else {
		dataS3.Buckets = listBuckets
	}
	//dataS3.Buckets = []storage.Bucket{{Name: "item 1"}, {Name: "item2"}, {Name: "item3"}}

	dataS3.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "s3.gohtml", dataS3)
}

func S3ObjectsHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	bucketname := request.Form["bucketname"]

	dataS3.Objects = []storage.BucketObj{}
	s3Client := storage.NewS3Client()
	listObjects, err := s3Client.ListObjects(bucketname[0])

	if err != nil {
		dataS3.Err = err.Error()
	} else {
		dataS3.Objects = listObjects
	}
	//dataS3.Objects = []storage.BucketObj{{Name: "obj 1"}, {Name: "obj"}, {Name: "obj"}}

	dataS3.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "s3.gohtml", dataS3)
}

func S3CreateHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	bucketname := request.Form["bucketname"]

	dataS3.Created = false
	s3Client := storage.NewS3Client()
	created, err := s3Client.CreateBucket(&bucketname[0])

	if err != nil {
		dataS3.Err = err.Error()
	} else {
		dataS3.Created = created
	}

	dataS3.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "s3.gohtml", dataS3)
}

func S3UploadHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	bucketname := request.Form["bucketname"]
	objectname := request.Form["objectname"]

	dataS3.Uploaded = false
	s3Client := storage.NewS3Client()
	uploaded, err := s3Client.Upload(&bucketname[0], &objectname[0])

	if err != nil {
		dataS3.Err = err.Error()
	} else {
		dataS3.Uploaded = uploaded
	}

	//dataS3.Uploaded = true
	dataS3.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "s3.gohtml", dataS3)
}
