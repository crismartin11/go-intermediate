package dynamodb

import (
	"fmt"

	"go-intermediate/pkg/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var (
	tableNameUsers string
)

func init() {
	tableNameUsers = "Users"
}

type DynamoDBClient struct{}

func NewDynamoDBClient() DynamoDBClient {
	return DynamoDBClient{}
}

type User struct {
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
}

func (db DynamoDBClient) ListUsers() ([]User, error) {
	service := credentials.GetClientDynamo()
	users := []User{}

	// Con la proyección obtengo el UserName e Email de cada elemento recuperado
	proj := expression.NamesList(expression.Name("UserName"), expression.Name("Email"))

	// Creo la expresión
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return users, fmt.Errorf("Error creando expresión: %s", err)
	}

	// Creo el objeto de parámetros de entrada
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableNameUsers),
	}

	// Invoco DynamoDB Query API
	result, err := service.Scan(params)
	if err != nil {
		return users, fmt.Errorf("Error al invocar API Query: %s", err)
	}

	// Recorro los items obtenidos
	for _, i := range result.Items {
		user := User{}

		err = dynamodbattribute.UnmarshalMap(i, &user) // Parseo y almaceno en user
		if err != nil {
			return users, fmt.Errorf("Error al parsear item: %s", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (db DynamoDBClient) Create(username string, email string) (bool, error) {
	service := credentials.GetClientDynamo()
	user := User{UserName: username, Email: email}

	// Parseo cada ítems de Go Types a DynamoDB attributes values
	us, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return false, fmt.Errorf("Error de parseo: %s", err)
	}

	// Creo el ítem a insertar en la tabla Users
	input := &dynamodb.PutItemInput{
		Item:      us,
		TableName: aws.String(tableNameUsers),
	}

	// Inserto
	_, err = service.PutItem(input)
	if err != nil {
		return false, fmt.Errorf("Error insertando usuario: %s", err)
	}

	return true, nil
}

func (db DynamoDBClient) Update(username string, email string) (bool, error) {
	service := credentials.GetClientDynamo()

	// Creo el ítem a actualizar en la tabla Users
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableNameUsers),
		Key: map[string]*dynamodb.AttributeValue{
			"UserName": {
				S: aws.String(username),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Email = :email"),
	}

	// Actualizo
	_, err := service.UpdateItem(input)
	if err != nil {
		return false, fmt.Errorf("Error actualizando usuario: %s", err)
	}

	return true, nil
}

func (db DynamoDBClient) Delete(username string) (bool, error) {
	service := credentials.GetClientDynamo()

	// Creo el ítem a eliminar en la tabla Users
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserName": {
				S: aws.String(username),
			},
		},
		TableName: aws.String(tableNameUsers),
	}

	// Elimino
	_, err := service.DeleteItem(input)
	if err != nil {
		return false, fmt.Errorf("Error eliminando usuario: %s", err)
	}

	return true, nil
}
