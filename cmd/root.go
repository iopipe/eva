package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eva",
	Short: "Event Ally, a cli for managing serverless events",
	Long: `Eva is a CLI application that enables developers
to work with events to store, replay, deliver,
and proxy. It is designed to work with event-driven
serverless systems.

Eva can generate events:
  `+"`"+`eva generate <event-type>`+"`"+`,
Consume and dispatch events as a daemon:
  `+"`"+`eva daemon <event-type>`+"`"+`,
Replay events and redispatch:
  `+"`"+`eva play <event-id>`+"`"+`,
Store invocation data for serverless functions:
  `+"`"+`eva invocations`+"`"+`,
and is your serverless event ally brought to you with <3

Read documentation online at https://iopipe.github.io/eva/`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eva.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".eva" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".eva")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
