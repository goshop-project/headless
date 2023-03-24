package main

import (
	"os"

	"darvaza.org/core"

	"goshop.dev/headless/pkg/mkcert"
)

// PrepareCA loads an existing CA or creates a new one if
// none was found
func PrepareCA(cfg *mkcert.Config) (*mkcert.CA, error) {
	ca, err := cfg.LoadCA()
	switch {
	case err == nil:
		// ready
		log.Info().Printf("CA loaded from %q", cfg.RootDir)
		return ca, nil
	case !os.IsNotExist(err):
		// load error
		fatal(err, "failed to load CA files")
		return nil, err
	default:
		// not found
		log.Warn().Printf("CA not found at %q", cfg.RootDir)

		// NewCA
		ca, err = cfg.NewCA()
		if err != nil {
			core.Wrap(err, "failed to create CA")
			return nil, err
		}

		log.Info().Println("new CA created")

		// write rootCA.pem
		fn, err := ca.WriteCertFile()
		if err != nil {
			core.Wrapf(err, "failed to write CA Certificate to %q", fn)
			return nil, err
		}
		log.Info().Printf("CA Certificate written to %q", fn)

		// write rootCA-key.pem
		fn, err = ca.WriteKeyFile()
		if err != nil {
			core.Wrapf(err, "failed to write CA Key to %q", fn)
			return nil, err
		}
		log.Info().Printf("CA Key written to %q", fn)

		return ca, nil
	}
}
