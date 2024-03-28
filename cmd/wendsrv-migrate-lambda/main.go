package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ag7if/wendover/internal/aws"
	"github.com/ag7if/wendover/internal/database"
	_ "github.com/ag7if/wendover/internal/logging"
)

func handler() error {
	migration, err := database.NewMigration()
	if err != nil {
		return errors.WithStack(err)
	}

	version, err := migration.Up()
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info().Uint("version", version).Msg("successfully migrated database")

	return nil
}

func main() {
	err := aws.SetupConfigFromParameterStore()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup config from AWS Parameter Store")
	}
	lambda.Start(handler)
}
