package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	"github.com/ag7if/wendover/internal/aws"
	_ "github.com/ag7if/wendover/internal/logging"
)

func handler(ctx context.Context) error {
	panic("implement me!")
}

func main() {
	err := aws.SetupConfigFromParameterStore()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup config from AWS Parameter Store")
	}
	lambda.Start(handler)
}
