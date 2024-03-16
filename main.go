/*
Copyright © 2024 Kapparina 116474667+Kapparina@users.noreply.github.com
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kapparina/Dotty/cmd"
	"github.com/Kapparina/Dotty/pkg/helpers/files"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var (
	dottyDir       *string      = new(string)
	envFile        *string      = new(string)
	dottyEnvConfig *viper.Viper = viper.New()
	// dottyRuntimeConfig *viper.Viper = new(viper.Viper)
)

func init() {
	notifyAttemptCreation := func(path string) {
		color.Set(color.FgYellow)
		fmt.Printf("`%s` does not exist - Attempting creation...\n", filepath.Clean(path))
		color.Unset()
	}
	homeDir, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		panic(fmt.Errorf("Fatal error getting homeDir directory: %w \n", homeDirErr))
	}
	*dottyDir = filepath.Join(homeDir, ".config", "dotty")
	defaultDottyConfigFile := filepath.Join(*dottyDir, "config.toml")

	*envFile = filepath.Join(*dottyDir, ".dotty")
	dottyEnvConfig.SetConfigType("env")
	dottyEnvConfig.SetConfigFile(*envFile)
	dottyEnvConfig.SetDefault("DOTTY_CONFIG_PATH", fmt.Sprintf("'%s'", defaultDottyConfigFile))

	if _, dottyDirErr := os.Stat(*dottyDir); os.IsNotExist(dottyDirErr) {
		notifyAttemptCreation(*dottyDir)
		newDottyDir()
	}
	if _, envFileErr := os.Stat(*envFile); os.IsNotExist(envFileErr) {
		notifyAttemptCreation(*envFile)
		newEnvFile()
		configWriteErr := dottyEnvConfig.WriteConfigAs(*envFile)
		if configWriteErr != nil {
			color.Set(color.FgHiRed)
			panic(fmt.Errorf("fatal error writing to .dotty file:\n%w\n", configWriteErr))
		}
	} else {
		envReadErr := dottyEnvConfig.ReadInConfig()
		if envReadErr != nil {
			color.Set(color.FgHiRed)
			panic(fmt.Errorf("fatal error reading .dotty file:\n%w\n", envReadErr))
		}
	}
}

func newDottyDir() {
	dirCreationErr := files.CreateDir(*dottyDir)
	if dirCreationErr != nil {
		color.Set(color.FgHiRed)
		panic(fmt.Errorf("Fatal error creating Dotty config directory:\n%w\n", dirCreationErr))
	}
	color.Set(color.FgHiGreen)
	fmt.Println("Dotty config directory created successfully!")
	color.Unset()
}

func newEnvFile() {
	fileCreationErr := files.CreateFile(*envFile)
	if fileCreationErr != nil {
		color.Set(color.FgHiRed)
		panic(fmt.Errorf("Fatal error creating .dotty file: \n%w\n", fileCreationErr))
	}
	color.Set(color.FgHiGreen)
	fmt.Println(".dotty file created successfully!")
	color.Unset()
}

func main() {
	cmd.Execute()
}
