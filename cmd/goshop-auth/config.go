package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"go.sancus.dev/config/yaml"

	"goshop.dev/headless/pkg/server"
)

// Config is the configuration structure of this microservice
type Config struct {
	Server server.Config
}

// ReadInFile loads the microservice configuration from a YAML file by name
func (cfg *Config) ReadInFile(filename string) error {
	return yaml.LoadFile(filename, cfg)
}

// WriteTo writes out the Config
func (cfg *Config) WriteTo(f io.Writer) (int64, error) {
	return yaml.WriteTo(f, cfg)
}

// SetDefaults fills any gap in the Config
func (cfg *Config) SetDefaults() error {
	return cfg.Server.SetDefaults()
}

// Command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump shows the loaded config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := cfg.WriteTo(os.Stdout); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
