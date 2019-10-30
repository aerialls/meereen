package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/aerialls/meereen/core"
	_ "github.com/aerialls/meereen/notifier"
	_ "github.com/aerialls/meereen/processor"
)

var (
	cfgFile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "meereen",
	Short: "Meereen is a lightweight monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		container := core.NewContainer()

		err := container.LoadConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		scheduler := core.NewScheduler(container)
		<-scheduler.Start()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	rootCmd.MarkFlagRequired("config")
}

func initConfig() {
	level := log.InfoLevel
	if verbose {
		level = log.DebugLevel
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
