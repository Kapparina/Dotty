/*
Copyright © 2024 Kapparina 116474667+Kapparina@users.noreply.github.com
*/

package cmd

import (
	"os"

	"github.com/Kapparina/Dotty/cmd/config"
	"github.com/Kapparina/Dotty/cmd/test"
	"github.com/spf13/cobra"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dotty.yaml)")
	rootCmd.AddCommand(test.Command)
	rootCmd.AddCommand(config.ConfigurationCmd)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotty",
	Short: "Dotty is a CLI tool for managing dotfiles",
	Long:  `Dotty is a CLI tool for managing dotfiles. It is designed to be simple and easy to use.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
