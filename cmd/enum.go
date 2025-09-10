package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/khuong02/kd-gen/config"
	"github.com/khuong02/kd-gen/pkg/enum"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputFile string
	pkgName    string
)

var enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "Commands for managing Enum",
}

var enumGenCmd = &cobra.Command{
	Use:          "gen (output-file)",
	Short:        "Generate enum from YAML config",
	Example:      "enum gen --package core --config ./example/enum/enum.yaml --output ./pkg/core/enum_gen.go",
	SilenceUsage: true,
	RunE:         enumGen,
	Args:         cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(enumCmd)
	enumCmd.AddCommand(enumGenCmd)

	enumGenCmd.Flags().StringVarP(&outputFile, "output", "o", "./pkg/core/enum_gen.go", "Output file path")
	enumGenCmd.Flags().StringVarP(&pkgName, "package", "p", "core", "Package name")
}

func enumGen(cmd *cobra.Command, args []string) error {
	if strings.TrimSpace(pkgName) == "" {
		return fmt.Errorf("package name is required")
	}

	cli := setupEnumCLI()

	cli.client.HeaderComment()
	var ImportName = enum.ImportName{
		"fmt": "fmt",
	}
	cli.client.ImportName(ImportName)

	f := cli.client.JenFile()

	for _, e := range cli.config.Enums {
		cli.client.Enum(e.Name, e.Type, e.Values, e.Methods)
	}

	// make sure the folder exists
	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if err := f.Save(outputFile); err != nil {
		return fmt.Errorf("save file: %w", err)
	}

	slog.Info("Generated enums to", "path", outputFile)
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
	w.client = enum.New(pkgName)
	return w
}
