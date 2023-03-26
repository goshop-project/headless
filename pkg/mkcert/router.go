package mkcert

import (
	"net/http"

	"darvaza.org/acmefy/pkg/acme"
	"darvaza.org/middleware"
)

// Router is the http.Handler of pkg/mkcert for a given ca.CA
type Router struct {
	*http.ServeMux

	ca *CA
}

// NewRouter creates a router
func NewRouter(m *CA) (http.Handler, error) {
	r := &Router{
		ServeMux: http.NewServeMux(),

		ca: m,
	}

	h := http.HandlerFunc(r.ca.ca.ServeCertificate)
	r.Handle("/cacert.pem", AcceptMiddleware(h,
		acme.ContentTypePEMCertChain,
		acme.ContentTypePEM))
	r.Handle("/cacert.cer", AcceptMiddleware(h,
		acme.ContentTypeDERCA))

	return r, nil
}

// AcceptMiddleware wraps the request so the handler receives a pre-decided
// media type based on the list of supported types
func AcceptMiddleware(h http.Handler, supported ...string) http.Handler {
	f := middleware.AcceptMiddleware(supported...)
	return f(h)
}
