package commands

import (
	"github.com/AnubhavUjjawal/ginango/utils"
	"github.com/spf13/cobra"
)

var projectName string
var startprojectCmd = &cobra.Command{
	Use:   "startproject",
	Short: "Starts a new ginango project",
	// TODO: Add flags that will allow to download a project template from the internet and use it.
	Run: func(cmd *cobra.Command, args []string) {
		utils.SugaredLogger.Infof("Bootstrapping a new project: %s", projectName)
		// TODO: Add command to copy files.
	},
}

func init() {
	startprojectCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "name of the project")
	startprojectCmd.MarkFlagRequired("project-name")
	rootCmd.AddCommand(startprojectCmd)
}
