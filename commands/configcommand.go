package commands

import (
	"github.com/ryanwalls/golang-seed/config"
	"github.com/ryanwalls/golang-seed/logger"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Example command that prints out config information",
	Run: func(cmd *cobra.Command, args []string) {
		// Example of using the logger and config packages
		logger.Log.Info("Printing configuration", "config", config.Viper.AllSettings())
	},
}
