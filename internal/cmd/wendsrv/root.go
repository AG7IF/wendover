package wendsrv

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var loglevel string

var rootCmd = &cobra.Command{
	Use:   "wendsrv",
	Short: "Backend service for the Wendover toolset",
	Long:  ``,
}

func Execute(version string) {
	rootCmd.Version = version

	err := rootCmd.Execute()
	if err != nil {
		log.Err(err).Msg("error when attempting to execute")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&loglevel, "loglevel", "info", "")
}
