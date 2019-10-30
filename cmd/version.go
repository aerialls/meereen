package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the Meereen version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Meeren", Version, Build)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
