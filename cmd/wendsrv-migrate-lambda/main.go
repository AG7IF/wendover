package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	"github.com/ag7if/wendover/internal/aws"
	"github.com/ag7if/wendover/internal/database"
	_ "github.com/ag7if/wendover/internal/logging"
)

func handler() (map[string]string, error) {
	closeLogs, err := aws.SetupConfigFromParameterStore()
	if err != nil {
		return map[string]string{
			"status":  "error",
			"message": "failed to set up config from param store",
			"error":   err.Error(),
		}, nil
	}

	migration, err := database.NewMigration()
	if err != nil {
		return map[string]string{
			"status":  "error",
			"message": "failed to set up config from param store",
			"error":   err.Error(),
		}, nil
	}

	version, err := migration.Up()
	if err != nil {
		return map[string]string{
			"status":  "error",
			"message": "failed to set up config from param store",
			"error":   err.Error(),
		}, nil
	}

	log.Info().Uint("version", version).Msg("successfully migrated database")
	err = closeLogs()

	return map[string]string{
		"status":  "success",
		"message": "successfully migrated database",
	}, nil
}

func main() {
	lambda.Start(handler)
}
