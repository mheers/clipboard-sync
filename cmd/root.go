package cmd

import (
	"github.com/mheers/clipboard-sync/helpers"
	"github.com/spf13/cobra"
)

var (
	// LogLevelFlag describes the verbosity of logs
	LogLevelFlag string
	// ConfigFileFlag holds the path to the config file
	ConfigFileFlag string

	// OutputFormatFlag can be json, yaml or table
	OutputFormatFlag string

	// // Config holds the read config
	// Config *config.Config

	rootCmd = &cobra.Command{
		Use:   "clipboardsyncclient",
		Short: "Clipboardsyncclient is a command line interface for the clipboard-sync client",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			helpers.PrintInfo()
			cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&LogLevelFlag, "log-level", "l", "error", "possible values are debug, error, fatal, panic, info, trace")
	rootCmd.PersistentFlags().StringVarP(&OutputFormatFlag, "output-format", "O", "table", "format [json|table|yaml|csv]")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(syncerclientCmd)
}
