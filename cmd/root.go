package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "0.2.3"
)

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "go-generate",
	Short: "Go Generate CLI Tool",
	Long:  `A CLI tool that generates Go enum types with useful helper methods from YAML configuration files.`,
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "config.yml", "Configuration file to use.")
	_ = viper.BindEnv("config")
	_ = viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
