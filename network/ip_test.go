package network

import (
	"net"
	"testing"
)

func TestCIDR(t *testing.T) {
	ip, ipnet, err := net.ParseCIDR("192.168.0.10")
	t.Logf("ip: %+v, ipnet: %+v, error: %v", ip, ipnet, err)
}
