package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"SecondLambda/pkg/handler"
)

func main() {
	handler := handler.New()
	lambda.Start(handler.Handle)
}

// set GOOS=linux
// set GOARCH=amd64
// set CGO_ENABLED=0
// go build -o main main.go		(genero el binario)
// %USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main	(comprimo en zip)

// Crear la función en aws
// Subir el zip asociada a la función creada
// Cambiar el nombre del handler en "Runtime settings" por main
