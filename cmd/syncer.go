package cmd

import (
	"github.com/mheers/clipboard-sync/config"
	"github.com/mheers/clipboard-sync/helpers"
	"github.com/mheers/clipboard-sync/syncer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	syncerclientCmd = &cobra.Command{
		Use:   "start",
		Short: "starts the syncerclient",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := config.GetConfig(false)

			// Set the log level
			helpers.SetLogLevel(LogLevelFlag)

			// Create the syncer
			logrus.Info("Creating the syncer")
			w, err := syncer.NewSyncer(config)
			if err != nil {
				return err
			}
			return w.Start()
		},
	}
)
