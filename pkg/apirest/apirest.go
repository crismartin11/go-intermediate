package apirest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-intermediate/pkg/dynamodb"
	"go-intermediate/pkg/storage"
)

func Start() {
	fmt.Println("Server started")
	http.HandleFunc("/buckets", bucketsHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}

func bucketsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s3Client := storage.NewS3Client()
	listBuckets, err := s3Client.ListBuckets()

	w.WriteHeader(http.StatusOK)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(listBuckets)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dynamoDBClient := dynamodb.NewDynamoDBClient()

	switch r.Method {
	case http.MethodGet:
		userName := r.URL.Query().Get("username")
		if userName != "" {
			usersHandlerGetByUserName(w, dynamoDBClient, userName)
		} else {
			usersHandlerGetList(w, dynamoDBClient)
		}

	case http.MethodPut:
		usersHandlerAddUser(w, r, dynamoDBClient)

	case http.MethodPost:
		usersHandlerAddUser(w, r, dynamoDBClient)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func usersHandlerGetList(w http.ResponseWriter, dynamoDBClient dynamodb.DynamoDBClient) {
	users, err := dynamoDBClient.ListUsers()

	w.WriteHeader(http.StatusOK)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func usersHandlerGetByUserName(w http.ResponseWriter, dynamoDBClient dynamodb.DynamoDBClient, userName string) {
	user, err := dynamoDBClient.UserByUserName(userName)

	w.WriteHeader(http.StatusOK)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(user)
}

func usersHandlerAddUser(w http.ResponseWriter, r *http.Request, dynamoDBClient dynamodb.DynamoDBClient) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	user := string(body)

	w.WriteHeader(http.StatusOK)
	us, err := dynamoDBClient.AddUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(us)
}

func usersHandlerUpdateUser(w http.ResponseWriter, r *http.Request, dynamoDBClient dynamodb.DynamoDBClient) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	user := string(body)

	w.WriteHeader(http.StatusOK)
	us, err := dynamoDBClient.UpdateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(us)
}
