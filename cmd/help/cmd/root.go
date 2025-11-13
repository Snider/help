package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Help module for the core framework",
	Long:  `Help is a module of the core framework that provides assistance and documentation functionality.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the help module!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
