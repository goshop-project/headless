package mkcert

import (
	"crypto/x509"
	"encoding/pem"
	"errors"

	"darvaza.org/acmefy/pkg/ca"
	"darvaza.org/darvaza/shared/x509utils"
)

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

// LoadCA creates a new CA using [Config.KeyFile] and [Config.CertFile]
func (cfg *Config) LoadCA() (*CA, error) {
	keyFileName, err := cfg.KeyFileName()
	if err != nil {
		return nil, err
	}

	certFileName, err := cfg.CertFileName()
	if err != nil {
		return nil, err
	}

	key, err := loadKey(keyFileName)
	if err != nil {
		return nil, err
	}

	certs, err := loadCerts(certFileName)
	if err != nil {
		return nil, err
	}

	p, err := cfg.Export().LoadCA(key, certs)
	if err != nil {
		return nil, err
	}

	m := &CA{
		ca:  p,
		cfg: *cfg,
	}

	return m, nil
}

func loadKey(filename string) (x509utils.PrivateKey, error) {
	var pk x509utils.PrivateKey
	var addErr error

	readErr := x509utils.ReadFilePEM(filename, func(_ string, b *pem.Block) bool {
		var term bool

		key, err := x509utils.BlockToPrivateKey(b)
		switch {
		case err != nil:
			// invalid block
			addErr = err
			term = true
		case pk != nil:
			// only one key is allowed
			addErr = errors.New("multiple keys found")
			term = true
		default:
			// store
			pk = key
		}

		return term
	})

	switch {
	case readErr != nil:
		return nil, readErr
	case addErr != nil:
		return nil, addErr
	default:
		return pk, nil
	}
}

func loadCerts(filename string) ([]*x509.Certificate, error) {
	var certs []*x509.Certificate
	var addErr error

	readErr := x509utils.ReadFilePEM(filename, func(_ string, b *pem.Block) bool {
		crt, err := x509utils.BlockToCertificate(b)
		if err != nil {
			// invalid block
			addErr = err
			return true
		}

		// store and continue
		certs = append(certs, crt)
		return false
	})

	switch {
	case readErr != nil:
		return nil, readErr
	case addErr != nil:
		return nil, addErr
	default:
		return certs, nil
	}
}
