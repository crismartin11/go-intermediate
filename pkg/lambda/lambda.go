package lambda

import (
	"fmt"

	"go-intermediate/pkg/credentials"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	//"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

/*var (
	tableNameUsers string
)*/

/*func init() {
	tableNameUsers = "Users"
}*/

type LambdaClient struct{}

func NewLambdaClient() LambdaClient {
	return LambdaClient{}
}

type Lambda struct {
	Name string `json:"Name"`
}

func (db LambdaClient) ListLambdas() ([]Lambda, error) {
	service := credentials.GetClientLambda()

	lambdas := []Lambda{}
	result, err := service.ListFunctions(nil)
	if err != nil {
		return lambdas, fmt.Errorf("Error creando expresión: %s", err)
	}

	for _, f := range result.Functions {
		lambda := Lambda{Name: aws.StringValue(f.FunctionName)}
		lambdas = append(lambdas, lambda)
	}

	// lambdas := []Lambda{{Name: "l1"}, {Name: "l2"}}
	return lambdas, nil
}
