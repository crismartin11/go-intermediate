package handlers

import (
	"net/http"

	"go-intermediate/pkg/lambda"
)

type customDataLambda struct {
	Lambdas []lambda.Lambda
	/*Created bool
	Updated bool
	Deleted bool*/
	Err string
	commonData
}

var (
	dataLambda = customDataLambda{
		Lambdas: []lambda.Lambda{},
		/*Created: false,
		Updated: false,
		Deleted: false,*/
		Err: "",
	}
)

func LambdaHandler(response http.ResponseWriter, request *http.Request) {
	dataLambda.commonData = GetCommonData(request)
	templates.ExecuteTemplate(response, "lambda.gohtml", dataLambda)
}

func LambdaListHandler(response http.ResponseWriter, request *http.Request) {
	lambdaClient := lambda.NewLambdaClient()
	listLambdas, err := lambdaClient.ListLambdas()
	if err != nil {
		dataLambda.Err = err.Error()
	} else {
		dataLambda.Lambdas = listLambdas
	}

	dataLambda.commonData = GetCommonData(request)
	//dataDb.Users = []dynamodb.User{{UserName: "item 1", Email: "email1"}, {UserName: "item2", Email: "email2"}, {UserName: "item3", Email: "email3"}}
	templates.ExecuteTemplate(response, "lambda.gohtml", dataLambda)
}
