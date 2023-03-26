package mkcert

import (
	"crypto/tls"
	"crypto/x509"

	"darvaza.org/darvaza/shared/storage"
)

var (
	_ storage.Store = (*CA)(nil)
)

// GetCAPool returns a reference to the Certificates Pool
func (m *CA) GetCAPool() *x509.CertPool {
	return m.ca.GetCAPool()
}

// GetCertificate returns the TLS Certificate that should be used
// for a given TLS request
func (m *CA) GetCertificate(chi *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return m.ca.GetCertificate(chi)
}
