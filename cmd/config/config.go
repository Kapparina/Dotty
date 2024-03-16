package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Reset bool

func init() {
	ConfigurationCmd.Flags().BoolVarP(
		&Reset,
		"reset",
		"r",
		false,
		"reset Dotty configuration",
	)
}

var ConfigurationCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage dotty configuration",
	Long:  `Manage dotty configuration. You can use this command to set, get, and reset dotty configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Reset {
			// reset dotty configuration
			fmt.Println("Resetting dotty configuration")
		}
	},
}
