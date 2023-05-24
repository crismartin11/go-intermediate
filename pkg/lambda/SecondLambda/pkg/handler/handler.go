package handler

import (
	"context"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

// Defino estructura del response en formato JSON
type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Handle(ctx context.Context, event MyEvent) (Response, error) {
	return Response{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       fmt.Sprintf("Hello %s!", event.Name),
		},
		nil
}
