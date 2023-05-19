package handlers

import (
	"net/http"

	"go-intermediate/pkg/dynamodb"
)

type customDataDb struct {
	Users   []dynamodb.User
	Created bool
	Updated bool
	Deleted bool
	Err     string
	commonData
}

var (
	dataDb = customDataDb{
		Users:   []dynamodb.User{},
		Created: false,
		Updated: false,
		Deleted: false,
		Err:     "",
	}
)

func DynamodbHandler(response http.ResponseWriter, request *http.Request) {
	dataDb.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "dynamodb.gohtml", dataDb)
}

func DynamoUsersHandler(response http.ResponseWriter, request *http.Request) {
	dynamoDBClient := dynamodb.NewDynamoDBClient()
	listUsers, err := dynamoDBClient.ListUsers()
	if err != nil {
		dataDb.Err = err.Error()
	} else {
		dataDb.Users = listUsers
	}

	dataDb.commonData = GetCommonData(request)
	//dataDb.Users = []dynamodb.User{{UserName: "item 1", Email: "email1"}, {UserName: "item2", Email: "email2"}, {UserName: "item3", Email: "email3"}}
	templates.ExecuteTemplate(response, "dynamodb.gohtml", dataDb)
}

func DynamoCreateHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form["username"]
	email := request.Form["email"]

	dynamoDBClient := dynamodb.NewDynamoDBClient()
	created, err := dynamoDBClient.Create(username[0], email[0])
	if err != nil {
		dataDb.Err = err.Error()
	} else {
		dataDb.Created = created
	}

	dataDb.commonData = GetCommonData(request)
	//dataDb.Created = true
	templates.ExecuteTemplate(response, "dynamodb.gohtml", dataDb)
}

func DynamoUpdateHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form["username"]
	email := request.Form["email"]

	dynamoDBClient := dynamodb.NewDynamoDBClient()
	updated, err := dynamoDBClient.Update(username[0], email[0])
	if err != nil {
		dataDb.Err = err.Error()
	} else {
		dataDb.Updated = updated
	}

	dataDb.commonData = GetCommonData(request)
	//dataDb.Updated = true
	templates.ExecuteTemplate(response, "dynamodb.gohtml", dataDb)
}

func DynamoDeleteHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form["username"]

	dynamoDBClient := dynamodb.NewDynamoDBClient()
	deleted, err := dynamoDBClient.Delete(username[0])
	if err != nil {
		dataDb.Err = err.Error()
	} else {
		dataDb.Deleted = deleted
	}

	dataDb.commonData = GetCommonData(request)
	//dataDb.Deleted = true
	templates.ExecuteTemplate(response, "dynamodb.gohtml", dataDb)
}
