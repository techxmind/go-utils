package access

import (
	"github.com/techxmind/logger"
)

import (
	"net"
	"strings"
)

type IPList struct {
	ips   map[string]bool
	cidrs []*net.IPNet
}

// Value格式: 111.10.11.3, 111.10.11.0/24
func NewIPList(vals ...string) *IPList {
	ips := make(map[string]bool)
	cidrs := make([]*net.IPNet, 0)
	for _, ip := range vals {
		if strings.Contains(ip, "/") {
			_, cidr, err := net.ParseCIDR(ip)
			if err != nil {
				logger.Errorw("ip list value invalid, ignore it", "ip", ip)
				continue
			}
			cidrs = append(cidrs, cidr)
		} else {
			ips[ip] = true
		}
	}

	return &IPList{
		ips:   ips,
		cidrs: cidrs,
	}
}

func (l *IPList) Contains(ip string) bool {
	if _, ok := l.ips[ip]; ok {
		return true
	}

	nIP := net.ParseIP(ip)

	if nIP == nil {
		return false
	}

	for _, cidr := range l.cidrs {
		if cidr.Contains(nIP) {
			return true
		}
	}

	return false
}
