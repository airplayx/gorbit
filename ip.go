package gorbit

import (
	"errors"
	"net"
)

func Ip2long(ip net.IP) (uint32, error) {
	if ip.To4() == nil {
		return 0, errors.New("not ipv4 address")
	}
	a := uint32(ip[12])
	b := uint32(ip[13])
	c := uint32(ip[14])
	d := uint32(ip[15])
	return uint32(a<<24 | b<<16 | c<<8 | d), nil
}

func Long2ip(ip uint32) (net.IP, error) {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	if net.IPv4(a, b, c, d).To4() == nil {
		return nil, errors.New("not ipv4 address")
	}
	return net.IPv4(a, b, c, d), nil
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return ips, nil
}
