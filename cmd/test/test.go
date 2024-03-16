package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "test",
	Short: "Test command",
	Long:  "Test command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test command")
	},
}
