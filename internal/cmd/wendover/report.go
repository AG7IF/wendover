package wendover

import (
	"strings"

	"github.com/ag7if/go-files"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ag7if/wendover/internal/controllers/reports"
)

var reportCmd = &cobra.Command{
	Use:   "report [report-type] [file-name]",
	Short: "Commands for ingesting reports",
	Long:  ``,
	Run:   runReportCmd,
	Args:  cobra.ExactArgs(2),
}

func runReportCmd(cmd *cobra.Command, args []string) {
	f, err := files.NewFile(args[1], &log.Logger)
	if err != nil {
		log.Fatal().Err(err).Str("file", args[1]).Msg("Failed to open report file.")
	}

	switch strings.ToLower(args[0]) {
	case "rz":
		err = reports.IngestRegistrationZoneReport(f)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to ingest report")
		}
	default:
		log.Fatal().Str("report_type", args[0]).Msg("Report type not recognized.")
	}
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
