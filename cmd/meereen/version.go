package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
	date    string
	builtBy string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the Meereen version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Meereen %s (commit %s, built at %s)", version, commit, date))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
