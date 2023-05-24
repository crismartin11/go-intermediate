package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Defino evento en formato JSON
type MyEvent struct {
	Name string `json:"name"`
}

// Defino manejador del evento. Retorna un string
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}

// set GOOS=linux
// set GOARCH=amd64
// set CGO_ENABLED=0
// go build -o main main.go		(genero el binario)
// %USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main	(comprimo en zip)

// Crear la función en aws
// Subir el zip asociada a la función creada
// Cambiar el nombre del handler en "Runtime settings" por main
