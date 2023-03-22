package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/darvaza-proxy/slog"
	"github.com/darvaza-proxy/slog/handlers/discard"

	"goshop.dev/headless/pkg/config"
)

// Config represents the generic configuration for goshop servers
type Config struct {
	Logger  slog.Logger     `toml:"-" valid:",required"`
	Context context.Context `toml:"-"`

	Name string `toml:"name" default:"nil.goshop.local" valid:"host,required"`

	Supervision SupervisionConfig `toml:"run"`
	Addresses   BindConfig        `toml:",omitempty"`
	TLS         TLSConfig         `toml:"tls"`
	HTTP        HTTPConfig        `toml:"http"`
}

// SupervisionConfig represents how graceful upgrades will operate
type SupervisionConfig struct {
	PIDFile         string        `toml:"pid_file"         default:"/tmp/tableflip.pid"`
	GracefulTimeout time.Duration `toml:"graceful_timeout" default:"5s"`
	HealthWait      time.Duration `toml:"health_wait"      default:"1s"`
}

// BindConfig refers to the IP addresses used by a GoShop Server
type BindConfig struct {
	Interfaces []string `toml:"interfaces"`
	Addresses  []string `toml:"addresses" valid:"ip"`
}

// TLSConfig contains information for setting up TLS clients and server
type TLSConfig struct {
	Key   string `toml:"key"`
	Cert  string `toml:"cert"`
	Roots string `toml:"caroot"`
}

// Validate tells if the configuration is worth a try
func (c *TLSConfig) Validate() error {
	if c.Key != "" || c.Cert != "" || c.Roots != "" {
		return nil
	}
	return errors.New("missing TLS information")
}

// HTTPConfig contains information for setting up the HTTP server
type HTTPConfig struct {
	Port              uint16        `toml:"port"                default:"8443" valid:"port"`
	InsecurePort      uint16        `toml:"insecure_port"       default:"8080" valid:"port"`
	EnableInsecure    bool          `toml:"enable_insecure"`
	MutualTLSOnly     bool          `toml:"mtls_only"`
	ReadTimeout       time.Duration `toml:"read_timeout"        default:"1s"`
	ReadHeaderTimeout time.Duration `toml:"read_header_timeout" default:"2s"`
	WriteTimeout      time.Duration `toml:"write_timeout"       default:"1s"`
	IdleTimeout       time.Duration `toml:"idle_timeout"        default:"30s"`
}

// SetDefaults fills the gaps in the Config
func (cfg *Config) SetDefaults() error {
	if cfg.Logger == nil {
		cfg.Logger = discard.New()
	}

	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	return config.SetDefaults(cfg)
}

// Validate tells if the configuration is worth a try
func (cfg *Config) Validate() error {
	err := config.Validate(cfg)
	if err != nil {
		return err
	}

	// context.Background is *0 so valid:",required" fails
	if cfg.Context == nil {
		return fmt.Errorf("%s: %s", "Context", "can not be nil")
	}

	return cfg.TLS.Validate()
}
