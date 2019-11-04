package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/aerialls/meereen/core"
	_ "github.com/aerialls/meereen/notifier"
	_ "github.com/aerialls/meereen/processor"
)

var (
	cfgFile string
	verbose bool
	logger  *logrus.Logger
)

var rootCmd = &cobra.Command{
	Use:   "meereen",
	Short: "Meereen is a lightweight monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		container := core.NewContainer(logger)

		err := container.LoadConfig(cfgFile)
		if err != nil {
			logger.Fatal(err)
		}

		scheduler := core.NewScheduler(container, logger)
		<-scheduler.Start()
	},
}

func init() {
	logger = logrus.New()

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	rootCmd.MarkFlagRequired("config")
}

func initConfig() {
	level := logrus.InfoLevel
	if verbose {
		level = logrus.DebugLevel
	}

	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
