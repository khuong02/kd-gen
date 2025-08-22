package cmd

import (
	"fmt"
	"github.com/khuong02/kd-gen/config"
	"github.com/khuong02/kd-gen/pkg/enum"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputFile string
)

var apiKeyCmd = &cobra.Command{
	Use:   "enum",
	Short: "Commands for managing Enum",
}

var apiKeyCreateCmd = &cobra.Command{
	Use:          "gen (output-file)",
	Short:        "Generate enum from YAML config",
	Example:      "enum gen --output ./pkg/core/enum_gen.go",
	SilenceUsage: true,
	RunE:         enumGen,
	Args:         cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(apiKeyCmd)
	apiKeyCmd.AddCommand(apiKeyCreateCmd)

	apiKeyCreateCmd.Flags().StringVarP(&outputFile, "output", "o", "./pkg/core/enum_gen.go", "Output file path")
}

func enumGen(cmd *cobra.Command, args []string) error {
	cli := setupEnumCLI()

	cli.client.HeaderComment()
	var ImportName = enum.ImportName{
		"fmt": "fmt",
	}
	cli.client.ImportName(ImportName)

	f := cli.client.JenFile()

	for _, e := range cli.config.Enums {
		cli.client.Enum(e.Name, e.Type, e.Values)
	}

	// make sure the folder exists
	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if err := f.Save(outputFile); err != nil {
		return fmt.Errorf("save file: %w", err)
	}

	fmt.Printf("✅ Generated enums to %s\n", outputFile)
	return nil
}

type EnumCLI struct {
	config config.Config
	client *enum.Enum
}

func setupEnumCLI() EnumCLI {
	var w EnumCLI
	configPath := viper.GetString("config")
	c := config.NewConfig(configPath)
	w.config = c
	w.client = enum.New()
	return w
}
