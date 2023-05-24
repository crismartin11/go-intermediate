package lambda_test

import (
	"context"
	"testing"

	"SecondLambda/pkg/handler"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	ctx := context.Background()

	t.Run("Test handle", func(t *testing.T) {
		hand := handler.New()
		response, err := hand.Handle(ctx, handler.MyEvent{Name: "cris"})

		assert.NoError(t, err)
		assert.Equal(t, "Hello cris!", response.Body)

	})

}
