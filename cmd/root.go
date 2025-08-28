package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version   = "0.1.0"
	commit    = "unset"
	buildTime = "unset"

	dryRun = (os.Getenv("DRY_RUN") == "true")
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

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				commit = fmt.Sprintf("%.6s", setting.Value)
			}
			if setting.Key == "vcs.time" {
				buildTime = setting.Value
			}
		}
	}
}
