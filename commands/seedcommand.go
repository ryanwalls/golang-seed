package commands

import (
	"github.com/spf13/cobra"
	"os"
)

// SeedCmd is the root command - update this command to match your needs
var SeedCmd = &cobra.Command{
	Use: "seed",
}

//Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	addCommands()
	if err := SeedCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}
}

func addCommands() {
	// Add more commands here that run your application
	SeedCmd.AddCommand(configCmd)
	SeedCmd.AddCommand(versionCmd)
}
