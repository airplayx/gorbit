package ip

import (
	"net"
	"testing"
)

func TestIp2long(t *testing.T) {
	t.Parallel()
	t.Log(Uint32(net.IPv4bcast))
	t.Log(Uint32(net.IPv4allsys))
	t.Log(Uint32(net.IPv4allrouter))
	t.Log(Uint32(net.IPv4zero))

	t.Log(Uint32(net.IPv6zero))
	t.Log(Uint32(net.IPv6unspecified))
	t.Log(Uint32(net.IPv6loopback))
	t.Log(Uint32(net.IPv6interfacelocalallnodes))
	t.Log(Uint32(net.IPv6linklocalallnodes))
	t.Log(Uint32(net.IPv6linklocalallrouters))
}

func TestLong2ip(t *testing.T) {
	t.Parallel()
	t.Log(Ipv4(4294967295))
	t.Log(Ipv4(3758096385))
	t.Log(Ipv4(3758096386))
	t.Log(Ipv4(0))
}

func TestLocalIPv4s(t *testing.T) {
	t.Parallel()
	t.Log(LocalIPv4s())
}
