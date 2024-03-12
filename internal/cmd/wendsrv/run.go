package wendsrv

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ag7if/wendover/internal/server"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the Wendover backend service",
	Long:  ``,
	Run:   runRun,
}

func runRun(cmd *cobra.Command, args []string) {
	router, err := server.NewServer(rootCmd.Version)
	if err != nil {
		log.Error().Err(err).Msg("failed to create wendsrv service")
		os.Exit(1)
	}

	err = router.Run("localhost:8080")
	if err != nil {
		log.Error().Err(err).Msg("failed to start wendsrv service")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
