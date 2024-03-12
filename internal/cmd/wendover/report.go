package wendover

import (
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Commands for ingesting reports",
	Long:  ``,
	Run:   runReportCmd,
}

func runReportCmd(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(reportCmd)
}
