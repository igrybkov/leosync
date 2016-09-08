package cmd

import (
	"fmt"
	"os"

	cmdns "github.com/igrybkov/cmdns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	filePath string
)

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "leosync",
	Short: "LeoSync is an app to sync words and translations to/from LinguaLeo",
	Long:  `LeoSync is an app to sync words and translations to/from LinguaLeo`,
}

var exportCmd = &cobra.Command{
	Use: "export",
}

var importCmd = &cobra.Command{
	Use: "import",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(importCmd)
	RootCmd.AddCommand(exportCmd)

	cobra.OnInitialize(initConfig)
	// Enable namespacing
	cmdns.Namespace(RootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.config/leosync.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".leosync")      // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config") // adding home/config directory as first search path
	viper.AddConfigPath("$HOME")         // adding home directory as first search path
	viper.AddConfigPath(".")             // name of config file (without extension)
	viper.AutomaticEnv()                 // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		cfgFile = "~/.config/.leosync.yaml"
	}
}
