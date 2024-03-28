package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) error {
	panic("implement me!")
}

func main() {
	lambda.Start(handler)
}
