package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/briancsparks/winmove/winmove"

  "github.com/spf13/cobra"
)

// toprimaryCmd represents the toprimary command
var toprimaryCmd = &cobra.Command{
	Use:   "toprimary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toprimary called")
    winmove.ToPrimary()
	},
}

func init() {
	rootCmd.AddCommand(toprimaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toprimaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toprimaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
