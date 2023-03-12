package main

import (
	"net/http"

	"go.sancus.dev/config/flags"
	"go.sancus.dev/config/flags/cobra"
	"goshop.dev/headless/pkg/server"
)

// Command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "runs authentication service",
	PreRun: func(cmd *cobra.Command, args []string) {
		flags.GetMapper(cmd.Flags()).Parse()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var r http.Handler
		var srv *server.Server

		// Logger
		cfg.Server.Logger = log

		// Prepare server
		srv, err := cfg.Server.New()
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
