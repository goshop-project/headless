package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"goshop.dev/headless/pkg/config"
	"goshop.dev/headless/pkg/server"
)

// Config is the configuration structure of this microservice
type Config struct {
	Server server.Config `toml:"server"`
}

// ReadInFile loads the microservice configuration from a YAML file by name
func (cfg *Config) ReadInFile(filename string) error {
	return config.LoadFile(filename, cfg)
}

// WriteTo writes out the Config
func (cfg *Config) WriteTo(f io.Writer) (int64, error) {
	return config.WriteTo(f, cfg)
}

// Prepare fills any gap in the Config and validate its content
func (cfg *Config) Prepare() error {
	return config.Prepare(cfg)
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
