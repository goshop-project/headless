package server

import (
	"fmt"

	"darvaza.org/core"
	"darvaza.org/gossipcache"
	"darvaza.org/gossipcache/transport"
	"darvaza.org/slog"
)

func (srv *Server) newGossipTransportConfig() *transport.Config {
	tc := &transport.Config{
		Logger:  srv.cfg.Logger,
		Context: srv.ctx,
		OnError: srv.gossipTransportError,
	}
	return tc
}

func (srv *Server) gossipTransportError(err error) {
	srv.cfg.Logger.Error().
		WithField(slog.ErrorFieldName, err).
		Printf("%T.%s: %s", srv, "gossipTransportError", err)
}

func (srv *Server) cacheBaseURL() string {
	var base string

	saddr := srv.cfg.Addresses.AdvertiseAddr
	addr, err := core.ParseAddr(saddr)

	if !addr.IsValid() {
		core.PanicWrapf(err, "invalid AdvertiseAddr %q", saddr)
	}

	if addr.Is6() {
		base = fmt.Sprintf("[%s]", addr.String())
	} else {
		base = addr.String()
	}

	if port := srv.cfg.HTTP.Port; port != 443 {
		base = fmt.Sprintf("%s:%v", base, port)
	}

	return "https://" + base
}

func (srv *Server) newGossipCacheConfig() *gossipcache.Config {
	gcc := &gossipcache.Config{
		Logger:          srv.cfg.Logger,
		Context:         srv.ctx,
		CacheBaseURL:    srv.cacheBaseURL(),
		CacheBasePath:   srv.cfg.Cache.Path,
		ClientTLSConfig: srv.getClientTLSConfig(),
		AcceptedRootCA:  srv.getAcceptedRootCA(),
	}

	return gcc
}

func (srv *Server) spawnGossipCacheServer() error {
	gtc := srv.newGossipTransportConfig()
	gt, err := transport.NewWithListeners(gtc, &srv.lsn.Gossip)
	if err != nil {
		return err
	}

	gcc := srv.newGossipCacheConfig()
	gc, err := gossipcache.NewGossipCacheCluster(gcc,
		gossipcache.WithDefaultLANConfig(),
		gossipcache.WithTransport(gt),
	)
	if err != nil {
		return err
	}

	srv.gc = gc
	return nil
}

func (*Server) cancelGossipCacheServer() {
}
