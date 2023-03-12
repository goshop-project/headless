package server

import (
	"net/http"

	"github.com/darvaza-proxy/darvaza/agent/httpserver"
)

func (srv *Server) newHTTPServerConfig() *httpserver.Config {
	hsc := &httpserver.Config{
		Logger:  srv.cfg.Logger,
		Context: srv.ctx,

		// Addresses
		Bind: httpserver.BindingConfig{
			Interfaces: srv.cfg.Addresses.Interfaces,
			Addresses:  srv.cfg.Addresses.Addresses,

			Port:          srv.cfg.HTTP.Port,
			PortInsecure:  srv.cfg.HTTP.InsecurePort,
			AllowInsecure: srv.cfg.HTTP.EnableInsecure,
		},

		// HTTP
		ReadTimeout:       srv.cfg.HTTP.ReadTimeout,
		ReadHeaderTimeout: srv.cfg.HTTP.ReadHeaderTimeout,
		WriteTimeout:      srv.cfg.HTTP.WriteTimeout,
		IdleTimeout:       srv.cfg.HTTP.IdleTimeout,
		HandleInsecure:    srv.cfg.HTTP.EnableInsecure,

		// TODO: ACME

		// TLS
		GetCertificate: srv.getGetCertificateForServer(),
		GetRootCAs:     srv.getRootCAsForServer(),
		GetClientCAs:   srv.getClientCAsForServer(),
	}

	return hsc
}

func (srv *Server) spawnHTTPServer(h http.Handler) error {
	srv.wg.Go(func() error {
		return srv.hs.Serve(h)
	})

	srv.wg.Go(func() error {
		<-srv.ctx.Done()
		srv.hs.Cancel()
		return srv.hs.Wait()
	})

	return nil
}
