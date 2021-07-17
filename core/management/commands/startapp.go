package commands

import (
	"github.com/AnubhavUjjawal/ginango/utils"
	"github.com/spf13/cobra"
)

var appName string
var startappCmd = &cobra.Command{
	Use:   "startapp",
	Short: "Starts a new ginango app",
	// TODO: Add flags that will allow to download a app template from the internet and use it.
	Run: func(cmd *cobra.Command, args []string) {
		utils.SugaredLogger.Infof("Creating a new app: %s", appName)
		// TODO: Add command to copy files.
	},
}

func init() {
	startappCmd.Flags().StringVarP(&appName, "app-name", "n", "", "name of the app")
	startappCmd.MarkFlagRequired("app-name")
	rootCmd.AddCommand(startappCmd)
}
