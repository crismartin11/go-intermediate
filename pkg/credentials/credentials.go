package credentials

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

func getSession() *session.Session {
	// Creo sesión tomando credenciales y región de ~/.aws/credentials y ~/.aws/config
	/*return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))*/

	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	aws_default_region := os.Getenv("AWS_DEFAULT_REGION")

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session, _ := session.NewSession(&aws.Config{
		Region:      aws.String(aws_default_region),
		Credentials: creds,
	})
	return session
}

func GetClientS3() *s3.S3 {
	// Instancio cliente de servicio S3
	return s3.New(getSession())
}

func GetClientLambda() *lambda.Lambda {
	// Create Lambda service client
	return lambda.New(getSession())
}

func GetClientDynamo() *dynamodb.DynamoDB {
	// Creo cliente de DynamoDB
	return dynamodb.New(getSession())
}
