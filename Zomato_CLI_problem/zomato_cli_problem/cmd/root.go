/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Cuisine string
var City string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Find restaurant",
	Short: "Enter your choices",
	Long:  "Enter your choices to help us find restaurant for you",
	Run: func(cmd *cobra.Command, args []string) {
		Cuisine, _ = cmd.Flags().GetString("cuisine")
		City, _ = cmd.Flags().GetString("city")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("city", "t", "", "Enter city here:")
	rootCmd.Flags().StringP("cuisine", "c", "", "Enter cuisine here:")
	// rootCmd.Flags().StringP("city", "t", "", "Enter city here:")
}
