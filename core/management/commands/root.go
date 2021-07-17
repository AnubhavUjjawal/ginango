package commands

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "Ginango: Web Framework inspired by Django.",
		Short: "Ginango is a highly opionionated rapid web development framework, written in go.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
