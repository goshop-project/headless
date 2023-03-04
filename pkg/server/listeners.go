package server

import (
	"net"
	"net/netip"

	"darvaza.org/core"
	"darvaza.org/darvaza/agent/httpserver"
	"darvaza.org/gossipcache/transport"
)

// Listeners contains all listeners of a Server
type Listeners struct {
	HTTP   httpserver.ServerListeners
	Gossip transport.Listeners
}

// Close closes all listeners of a Server
func (sl *Listeners) Close() error {
	err1 := sl.HTTP.Close()
	err2 := sl.Gossip.Close()

	if err1 != nil {
		return err1
	}
	return err2
}

// Port tells the port used for HTTPS
func (sl *Listeners) Port() uint16 {
	if len(sl.HTTP.Secure) > 0 {
		addr := sl.HTTP.Secure[0].Addr().(*net.TCPAddr)
		return uint16(addr.Port)
	}
	return 0
}

// InsecurePort tells the port used for HTTP
func (sl *Listeners) InsecurePort() uint16 {
	if len(sl.HTTP.Insecure) > 0 {
		addr := sl.HTTP.Insecure[0].Addr().(*net.TCPAddr)
		return uint16(addr.Port)
	}
	return 0
}

// GossipPort tells the port used for memberlist
func (sl *Listeners) GossipPort() uint16 {
	if len(sl.Gossip.TCP) > 0 {
		addr := sl.Gossip.TCP[0].Addr().(*net.TCPAddr)
		return uint16(addr.Port)
	}
	return 0
}

// IPAddresses tells the IP Addresses used by this Server
func (sl *Listeners) IPAddresses() ([]netip.Addr, error) {
	addrs, err := sl.HTTP.IPAddresses()
	out := make([]netip.Addr, 0, len(addrs))

	for _, ip := range addrs {
		if addr, ok := netip.AddrFromSlice(ip); ok {
			if addr.IsValid() {
				out = append(out, addr.Unmap())
			}
		}
	}
	return out, err
}

// NetIPAddresses tells the net.IP Addresses used by this Server
func (sl *Listeners) NetIPAddresses() ([]net.IP, error) {
	return sl.HTTP.IPAddresses()
}

// StringIPAddresses tells the IP Addresses used by this Server, as strings
func (sl *Listeners) StringIPAddresses() ([]string, error) {
	return sl.HTTP.StringIPAddresses()
}

// AdvertiseIPAddress tells the first IP Address in the set
func (sl *Listeners) AdvertiseIPAddress() (netip.Addr, bool) {
	if len(sl.HTTP.Secure) > 0 {
		if addrport, ok := core.AddrPort(sl.HTTP.Secure[0].Addr()); ok {
			addr := addrport.Addr()
			if addr.IsUnspecified() {
				addr = guessIPAddress()
			}
			return addr, true
		}
	}
	return netip.Addr{}, false
}

func guessIPAddress() netip.Addr {
	// let's try to skip the loopback
	ifaces, err := core.GetInterfacesNames("lo")
	if err == nil {
		addr := getFirstIPAddress(ifaces...)
		if addr.IsValid() {
			return addr
		}
	}

	// any then
	addr := getFirstIPAddress()
	if addr.IsValid() {
		return addr
	}

	// really? none?
	return netip.IPv4Unspecified()
}

func getFirstIPAddress(ifaces ...string) netip.Addr {
	addrs, _ := core.GetIPAddresses(ifaces...)
	if len(addrs) > 0 {
		return addrs[0]
	}

	return netip.Addr{}
}
