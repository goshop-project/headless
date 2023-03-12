package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/darvaza-proxy/slog"
	"github.com/darvaza-proxy/slog/handlers/discard"
	"go.sancus.dev/config"
)

// Config represents the generic configuration for goshop servers
type Config struct {
	Logger  slog.Logger     `yaml:"-" valid:",required"`
	Context context.Context `yaml:"-"`

	Name string `default:"nil.goshop.local" valid:"host,required"`

	Supervision SupervisionConfig `yaml:"run"`
	Addresses   BindConfig        `yaml:",omitempty"`
	TLS         TLSConfig
	HTTP        HTTPConfig
}

// SupervisionConfig represents how graceful upgrades will operate
type SupervisionConfig struct {
	PIDFile         string        `yaml:"pid-file"         default:"/tmp/tableflip.pid"`
	GracefulTimeout time.Duration `yaml:"graceful-timeout" default:"5s"`
	HealthWait      time.Duration `yaml:"health-wait"      default:"1s"`
}

// BindConfig refers to the IP addresses used by a GoShop Server
type BindConfig struct {
	Interfaces []string `yaml:"interfaces"`
	Addresses  []string `yaml:"addresses" valid:"ip"`
}

// TLSConfig contains information for setting up TLS clients and server
type TLSConfig struct {
	Key   string
	Cert  string
	Roots string
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
	Port              uint16        `yaml:"port"                default:"8443" valid:"port"`
	InsecurePort      uint16        `yaml:"insecure-port"       default:"8080" valid:"port"`
	EnableInsecure    bool          `yaml:"enable-insecure"`
	MutualTLSOnly     bool          `yaml:"mtls-only"`
	ReadTimeout       time.Duration `yaml:"read-timeout"        default:"1s"`
	ReadHeaderTimeout time.Duration `yaml:"read-header-timeout" default:"2s"`
	WriteTimeout      time.Duration `yaml:"write-timeout"       default:"1s"`
	IdleTimeout       time.Duration `yaml:"idle-timeout"        default:"30s"`
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
	_, err := config.Validate(cfg)
	if err != nil {
		return err
	}

	// context.Background is *0 so valid:",required" fails
	if cfg.Context == nil {
		return fmt.Errorf("%s: %s", "Context", "can not be nil")
	}

	return cfg.TLS.Validate()
}
