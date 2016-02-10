package commands

import (
	"github.com/ryanwalls/golang-seed/config"

	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang seed version: " + config.Viper.GetString("VERSION_TAG"))
	},
}
