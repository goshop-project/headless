package server

import (
	"crypto/tls"
	"crypto/x509"
)

func (srv *Server) getGetCertificateForServer() func(*tls.ClientHelloInfo) (*tls.Certificate,
	error) {
	return srv.tls.GetCertificate
}

func (*Server) getRootCAsForServer() func() *x509.CertPool {
	return nil
}

func (srv *Server) getClientCAsForServer() func() *x509.CertPool {
	if srv.cfg.HTTP.MutualTLSOnly {
		return srv.tls.GetCAPool
	}
	return nil
}

func (*Server) getClientTLSConfig() *tls.Config {
	return &tls.Config{}
}

func (*Server) getAcceptedRootCA() *x509.CertPool {
	return x509.NewCertPool()
}
