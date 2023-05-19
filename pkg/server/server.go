package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"go-intermediate/pkg/okta/utils"
	"go-intermediate/pkg/server/handlers"
)

var (
	templates *template.Template
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	templates = template.Must(template.ParseGlob(wd + "/web/templates/*"))
	utils.ParseEnvironment()
}

func Start() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc(utils.CALLBACK_URL, handlers.AuthCodeCallbackHandler) // A donde retorna okta
	http.HandleFunc("/logout", handlers.LogoutHandler)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/s3", handlers.S3Handler)
	http.HandleFunc("/dynamodb", handlers.DynamodbHandler)
	http.HandleFunc("/lambda", handlers.LambdaHandler)

	http.HandleFunc("/s3Buckets", handlers.S3BucketsHandler)
	http.HandleFunc("/s3Objects", handlers.S3ObjectsHandler)
	http.HandleFunc("/s3Create", handlers.S3CreateHandler)
	http.HandleFunc("/s3Upload", handlers.S3UploadHandler)

	http.HandleFunc("/dynamoUsers", handlers.DynamoUsersHandler)
	http.HandleFunc("/dynamoCreate", handlers.DynamoCreateHandler)
	http.HandleFunc("/dynamoUpdate", handlers.DynamoUpdateHandler)
	http.HandleFunc("/dynamoDelete", handlers.DynamoDeleteHandler)

	http.HandleFunc("/lambdaList", handlers.LambdaListHandler)

	fmt.Println("Server starting at localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Printf("Error %v", err)
	if err != nil {
		fmt.Printf("The HTTP server failed to start: %s", err)
		log.Fatalf("Error%s", err)
		os.Exit(1)
	}
}

func homeHandler(response http.ResponseWriter, request *http.Request) {
	data := handlers.GetCommonData(request)
	templates.ExecuteTemplate(response, "home.gohtml", data)
}
