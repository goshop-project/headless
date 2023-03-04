// Package server provides the skeleton for goshop services
package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"darvaza.org/acmefy/pkg/client/magic"
	"darvaza.org/core"
	"darvaza.org/darvaza/agent/httpserver"
	"darvaza.org/darvaza/shared/storage"
	"darvaza.org/darvaza/shared/storage/simple"
	"darvaza.org/gossipcache"
)

// A Server to run a node of a GoShop microservice
type Server struct {
	cfg       Config
	ctx       context.Context
	cancel    context.CancelFunc
	cancelled atomic.Bool
	err       atomic.Value
	wg        core.WaitGroup

	lsn *Listeners
	tls storage.Store
	hs  *httpserver.Server
	gc  *gossipcache.GossipCache
}

// New creates a new server using the given config
func (cfg *Config) New() (*Server, error) {
	return New(cfg)
}

// NewWithStore creates a new server using the given config and
// a prebuilt tls Store
func (cfg *Config) NewWithStore(s storage.Store) (*Server, error) {
	cfg.Store = s
	return New(cfg)
}

// New creates a new server using the given config
func New(cfg *Config) (*Server, error) {
	var err error

	if cfg == nil {
		cfg = &Config{}

		if err = cfg.SetDefaults(); err != nil {
			return nil, err
		}
	}

	if err = cfg.Validate(); err != nil {
		return nil, err
	}

	// TLS
	s := cfg.Store
	if s == nil {
		s, err = newTLSStore(cfg)
		if err != nil {
			return nil, err
		}
	}

	return cfg.newServer(s)
}

func newTLSStore(cfg *Config) (storage.Store, error) {
	if cfg.TLS.ACME == "" {
		sc := &simple.Config{
			Logger: cfg.Logger,
		}

		s, err := sc.New(cfg.TLS.Key,
			cfg.TLS.Cert,
			cfg.TLS.Roots)

		return s, err
	}

	sc := &magic.Config{
		URL:    cfg.TLS.ACME,
		Logger: cfg.Logger,
		Key:    cfg.TLS.Key,
		Cert:   cfg.TLS.Cert,
		Roots:  cfg.TLS.Roots,
	}

	s, err := sc.New()
	if err == nil {
		err = s.Start()
	}

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (cfg *Config) newServer(s storage.Store) (*Server, error) {
	ctx, cancel := context.WithCancel(cfg.Context)

	srv := &Server{
		cfg:    *cfg,
		cancel: cancel,
		ctx:    ctx,
		tls:    s,
	}

	return srv, nil
}

// Shutdown initiates a Shutdown with optional fatal timeout,
// and waits until all workers are done
func (srv *Server) Shutdown(timeout time.Duration) error {
	var ok atomic.Bool

	// once srv.Wait() finishes, we are done
	defer ok.Store(true)

	srv.tryCancel(nil)

	if timeout > 0 {
		time.AfterFunc(timeout, func() {
			if !ok.Load() {
				srv.fatal(nil).Print("graceful shutdown timed out")
			}
		})
	}

	return srv.Wait()
}

// Cancel initiates a shutdown of all workers
func (srv *Server) Cancel() {
	srv.tryCancel(nil)
}

// Fail initiates a shutdown with a reason
func (srv *Server) Fail(err error) {
	srv.tryCancel(err)
}

func (srv *Server) tryCancel(err error) {
	// once
	if srv.cancelled.CompareAndSwap(false, true) {
		if err != nil {
			srv.err.CompareAndSwap(nil, err)
		}
		srv.cancel()
	}
}

// Cancelled tells if the server has been cancelled
func (srv *Server) Cancelled() bool {
	return srv.cancelled.Load()
}

// Err returns the reason of the shutdown, if any
func (srv *Server) Err() error {
	if err, ok := srv.err.Load().(error); ok {
		return err
	} else if srv.Cancelled() {
		return os.ErrClosed
	} else {
		return nil
	}
}

// Wait blocks until all workers have exited
func (srv *Server) Wait() error {
	srv.wg.Wait()
	return srv.Err()
}

// revive:disable:cognitive-complexity

// Spawn the workers
func (srv *Server) Spawn(h http.Handler, healthy time.Duration) error {
	// revive:enable:cognitive-complexity
	var ok bool

	defer func() {
		if !ok {
			srv.Cancel()
		}
	}()

	if srv.lsn == nil {
		return errors.New("forgot to Listen?")
	}

	defer func() {
		if !ok {
			_ = srv.lsn.Close()
		}
	}()

	// GossipCache
	if err := srv.spawnGossipCacheServer(); err != nil {
		return err
	}

	// HTTP
	if err := srv.spawnHTTPServer(h); err != nil {
		srv.cancelGossipCacheServer()
		return err
	}

	if healthy > 0 {
		time.Sleep(healthy)
	}

	ok = true
	return srv.Err()
}
