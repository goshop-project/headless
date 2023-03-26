package main

import (
	"go.sancus.dev/config/flags"
	"go.sancus.dev/config/flags/cobra"

	"goshop.dev/headless/pkg/mkcert"
)

// Command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "runs CA service",
	PreRun: func(cmd *cobra.Command, args []string) {
		flags.GetMapper(cmd.Flags()).Parse()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Logger
		cfg.Server.Logger = log

		// Prepare CA
		ca, err := PrepareCA(&cfg.CA)
		if err != nil {
			return err
		}

		r, err := mkcert.NewRouter(ca)
		if err != nil {
			return err
		}

		// Prepare server
		srv, err := cfg.Server.NewWithStore(ca)
		if err != nil {
			return err
		}

		return srv.ListenAndServe(r)
	},
}

// Flags
func init() {
	cobra.NewMapper(serveCmd.Flags()).
		VarP(&cfg.Server.HTTP.Port, "port", 'p', "HTTPS Port").
		Var(&cfg.Server.Supervision.PIDFile, "pid", "Path to PID file").
		VarP(&cfg.Server.Supervision.GracefulTimeout, "graceful", 't',
			"Maximum time to wait for in-flight requests")

	rootCmd.AddCommand(serveCmd)
}
