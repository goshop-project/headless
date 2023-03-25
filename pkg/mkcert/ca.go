package mkcert

import "darvaza.org/acmefy/pkg/ca"

// CA is a private PKI for mTLS
type CA struct {
	ca  *ca.CA
	cfg Config
}

// NewCA generates a new CA using cfg.Issuer information
func (cfg *Config) NewCA() (*CA, error) {
	p, err := cfg.Export().NewCA(&cfg.Issuer)
	if err != nil {
		return nil, err
	}

	m := &CA{
		ca:  p,
		cfg: *cfg,
	}

	return m, nil
}
