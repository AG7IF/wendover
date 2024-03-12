package wendover

import (
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Commands for managing activities",
	Long:  ``,
	Run:   runActivityCmd,
}

func runActivityCmd(cmd *cobra.Command, args []string) {

}

var createActivityCmd = &cobra.Command{
	Use:   "create [activity_key] [activity_name] [activity_director_username]",
	Short: "Create a new activity",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run:   runCreateActivityCmd,
}

func runCreateActivityCmd(cmd *cobra.Command, args []string) {

}

var deleteActivityCmd = &cobra.Command{
	Use:   "delete [activity_key]",
	Short: "Delete all data related to the given activity from the database.",
	Long:  ``,
	Run:   runDeleteActivityCmd,
}

func runDeleteActivityCmd(cmd *cobra.Command, args []string) {

}

func init() {
	activityCmd.AddCommand(createActivityCmd)
	activityCmd.AddCommand(deleteActivityCmd)

	rootCmd.AddCommand(activityCmd)
}
