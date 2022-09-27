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
var HasTableBooking bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Find restaurant",
	Short: "Enter your choices",
	Long:  "Enter your choices to help us find restaurant for you",
	Run: func(cmd *cobra.Command, args []string) {
		Cuisine, _ = cmd.Flags().GetString("cuisine")
		City, _ = cmd.Flags().GetString("city")
		HasTableBooking, _ = cmd.Flags().GetBool("hasTableBooking")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("city", "t", "default", "Enter city here:")
	rootCmd.Flags().StringP("cuisine", "c", "default", "Enter cuisine here:")
	rootCmd.Flags().BoolP("hasTableBooking", "r", false, "Do you want to make reservations:")
}
