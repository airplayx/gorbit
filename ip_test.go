package gorbit

import (
	"net"
	"testing"
)

func TestIp2long(t *testing.T) {
	t.Parallel()
	t.Log(Ip2long(net.IPv4bcast))
	t.Log(Ip2long(net.IPv4allsys))
	t.Log(Ip2long(net.IPv4allrouter))
	t.Log(Ip2long(net.IPv4zero))

	t.Log(Ip2long(net.IPv6zero))
	t.Log(Ip2long(net.IPv6unspecified))
	t.Log(Ip2long(net.IPv6loopback))
	t.Log(Ip2long(net.IPv6interfacelocalallnodes))
	t.Log(Ip2long(net.IPv6linklocalallnodes))
	t.Log(Ip2long(net.IPv6linklocalallrouters))
}

func TestLong2ip(t *testing.T) {
	t.Parallel()
	t.Log(Long2ip(4294967295))
	t.Log(Long2ip(3758096385))
	t.Log(Long2ip(3758096386))
	t.Log(Long2ip(0))
}
