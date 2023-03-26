package mkcert

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"darvaza.org/acmefy/pkg/ca"
	"darvaza.org/core"

	"goshop.dev/headless/pkg/config"
)

const (
	// DefaultKeyFilePermissions are the permissions used when creating the KeyFile
	DefaultKeyFilePermissions os.FileMode = 0600
	// DefaultCertFilePermissions are the permissions used when creating the CertFile
	DefaultCertFilePermissions os.FileMode = 0644
	// DefaultDirPermissions are the permissions used when creating the RootDir
	DefaultDirPermissions os.FileMode = 0700
)

// Config is the configuration of a CA
type Config struct {
	// RootDir is the directory where key and cert are securely stored
	RootDir string `toml:"caroot_dir" default:"~/.local/share/mkcert"`
	// KeyFile is the filename of the CA key within the RootDir
	KeyFile string `toml:"key_file"   default:"rootCA-key.pem"`
	// CertFile is the filename of the CA certificate within the RootDir
	CertFile string `toml:"cert_file" default:"rootCA.pem"`

	// Certificates include infomation for the issued certificates
	Certificates ca.TemplateConfig `toml:"certificates,omitempty"`
	// Issuer include infomation for creating the CA
	Issuer ca.TemplateConfig `toml:"issuer,omitempty"`

	// KeyAlgorithm indicates the KeyAlgorigthm  to be used.
	// RSA, ECDSA, or ED25519. case-insensitive.
	KeyAlgorithm string `toml:"algorithm,omitempty"`
}

// Export converts an annotated [Config] into a [ca.Config]
func (cfg *Config) Export() *ca.Config {
	var algo ca.KeyAlgorithm

	switch cfg.KeyAlgorithm {
	case "ed25519":
		algo = ca.KeyAlgorithmED25519
	case "ecdsa":
		algo = ca.KeyAlgorithmECDSA
	default:
		algo = ca.KeyAlgorithmRSA
	}

	return &ca.Config{
		KeyAlgorithm: algo,
		Template:     cfg.Certificates,
	}
}

// SetDefaults fills the gaps in the [Config]
func (cfg *Config) SetDefaults() error {
	err := config.SetDefaults(cfg)
	if err != nil {
		return err
	}

	// KeyAlgorithm
	algo := strings.ToLower(cfg.KeyAlgorithm)
	switch {
	case algo == "":
		// default
		algo = "rsa"
	case !core.SliceContains([]string{
		"ed25519", "ecdsa", "rsa"}, algo):
		return fmt.Errorf("%q: unrecognised algorithm", algo)
	}
	cfg.KeyAlgorithm = algo

	tcI.SetDefaults(&cfg.Issuer)
	tcC.SetDefaults(&cfg.Certificates)

	// expand ~
	if strings.HasPrefix(cfg.RootDir, "~/") {
		u, err := user.Current()
		if err == nil {
			homedir := u.HomeDir
			cfg.RootDir = homedir + cfg.RootDir[1:]
		}
	}
	return nil
}

// KeyFileName returns the full path to the [Config.KeyFile]
func (cfg *Config) KeyFileName() (string, error) {
	fn := filepath.Join(cfg.RootDir, cfg.KeyFile)
	return fn, nil
}

// CertFileName returns the full path to the [Config.CertFile]
func (cfg *Config) CertFileName() (string, error) {
	fn := filepath.Join(cfg.RootDir, cfg.CertFile)
	return fn, nil
}

// MkRootDir creates the [Config.RootDir] if it doesn't exist already
func (cfg *Config) MkRootDir() error {
	return os.MkdirAll(cfg.RootDir, DefaultDirPermissions)
}

var userAndHostname string
var tcI, tcC ca.TemplateConfig

func getUserAndHostname() string {
	hostname, _ := os.Hostname()

	if u, err := user.Current(); err == nil {
		if s := generateUserAndHostname(u, hostname); s != "" {
			return s
		}
	}

	if hostname != "" {
		return hostname
	}

	return "undetermined"
}

func generateUserAndHostname(u *user.User, hostname string) string {
	var s []string

	// username@hostname (Name)
	if u.Username != "" {
		username := u.Username
		if hostname != "" {
			username += "@" + hostname
		}
		s = append(s, username)
	}

	if u.Name != "" && u.Name != u.Username {
		name := "(" + u.Name + ")"
		s = append(s, name)
	}

	return strings.Join(s, " ")
}

func init() {
	userAndHostname = getUserAndHostname()

	tcI = ca.TemplateConfig{
		O:  "mkcert development CA",
		OU: userAndHostname,
		CN: userAndHostname,

		Duration: ca.DefaultCADuration,
	}
	tcC = ca.TemplateConfig{
		O:  "mkcert development",
		OU: userAndHostname,

		Duration: ca.DefaultCertificateDuration,
	}
}
