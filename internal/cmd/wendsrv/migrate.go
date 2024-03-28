package wendsrv

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ag7if/wendover/internal/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database of the Wendover backend service to the latest version.",
	Long:  ``,
	Run:   runMigrate,
}

func runMigrate(cmd *cobra.Command, args []string) {
	migration, err := database.NewMigration()
	if err != nil {
		log.Error().Err(err).Msg("failed to generate migration for database")
		os.Exit(1)
	}

	version, err := migration.Up()
	if err != nil {
		log.Error().Err(err).Msg("failed to migrate database")
		os.Exit(1)
	}

	log.Info().Uint("version", version).Msg("successfully migrated database")
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
