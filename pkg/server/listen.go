package server

import (
	"syscall"

	"darvaza.org/core"
	"darvaza.org/darvaza/shared/net/bind"
)

// Listen listens to all needed ports
func (srv *Server) Listen() error {
	lc := bind.NewListenConfig(srv.ctx, 0)
	return srv.ListenWithListener(lc)
}

// ListenWithUpgrader listens to all needed ports using a ListenUpgrader
// like tableflip
func (srv *Server) ListenWithUpgrader(upg bind.Upgrader) error {
	lc := bind.NewListenConfig(srv.ctx, 0)
	return srv.ListenWithListener(lc.WithUpgrader(upg))
}

// revive:disable:cognitive-complexity

// ListenWithListener listens to all needed ports using a net.ListenerConfig
// context
func (srv *Server) ListenWithListener(lc bind.TCPUDPListener) error {
	// revive:enable:cognitive-complexity
	var sl Listeners
	var ok bool

	if srv.lsn != nil {
		return syscall.EBUSY
	}

	defer func() {
		if !ok {
			_ = sl.Close()
		}
	}()

	cfg := &srv.cfg

	// HTTPS
	bc := &bind.Config{
		Interfaces:  cfg.Addresses.Interfaces,
		Addresses:   cfg.Addresses.Addresses,
		Port:        cfg.HTTP.Port,
		DefaultPort: 8443,
	}
	bc.UseListener(lc)

	tcpLsn, udpLsn, err := bind.Bind(bc)
	if err != nil {
		return err
	}

	sl.HTTP.Secure = tcpLsn
	sl.HTTP.Quic = udpLsn

	// update bc.Addresses
	bc.RefreshFromTCPListeners(tcpLsn)

	// memberlist
	bc.Port = cfg.Cache.Port
	bc.DefaultPort = 7946

	tcpLsn, udpLsn, err = bind.Bind(bc)
	if err != nil {
		return err
	}

	sl.Gossip.TCP = tcpLsn
	sl.Gossip.UDP = udpLsn

	// HTTP
	if srv.cfg.HTTP.EnableInsecure {
		bc.Port = srv.cfg.HTTP.InsecurePort
		bc.DefaultPort = 8080
		bc.OnlyTCP = true

		tcpLsn, _, err = bind.Bind(bc)
		if err != nil {
			return err
		}

		sl.HTTP.Insecure = tcpLsn
	}

	// done
	cfg.refresh(bc, &sl)
	ok = true
	srv.lsn = &sl
	return nil
}

func (cfg *Config) refresh(bc *bind.Config, sl *Listeners) {
	// IP Addresses
	cfg.Addresses.Interfaces = bc.Interfaces
	cfg.Addresses.Addresses = bc.Addresses

	if addr, _ := sl.AdvertiseIPAddress(); addr.IsValid() {
		cfg.Addresses.AdvertiseAddr = addr.String()
	} else {
		core.Panic("unreachable")
	}

	// HTTP
	cfg.HTTP.Port = sl.Port()
	if port := sl.InsecurePort(); port > 0 {
		cfg.HTTP.InsecurePort = port
		cfg.HTTP.EnableInsecure = true
	} else {
		cfg.HTTP.InsecurePort = 0
		cfg.HTTP.EnableInsecure = false
	}

	// gossipcache
	cfg.Cache.Port = sl.GossipPort()
}
