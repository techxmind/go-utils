package access

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

var (
	ErrInvalidIP = errors.New("IP address is not valid")
	_cidrs       []*net.IPNet
)

func init() {
	maxCidrBlocks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	_cidrs = make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		_cidrs[i] = cidr
	}
}

// IsIntranetIP reports whether ip is intranet ip.
//
// Invalid ip return false.
func IsIntranetIPAddress(ipStr string) bool {
	return IsIntranetIP(net.ParseIP(ipStr))
}

// IsIntranetIP reports whether ip is intranet ip.
//
// Invalid ip return false.
func IsIntranetIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	for i := range _cidrs {
		if _cidrs[i].Contains(ip) {
			return true
		}
	}

	return false
}

// IsPublicIP reports whether ip is public network address.
//
// Invalid ip return false.
func IsPublicIPAddress(ipStr string) bool {
	return IsPublicIP(net.ParseIP(ipStr))
}

// IsPublicIP reports whether ip is public network address.
//
// Invalid ip return false.
func IsPublicIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	return ip.IsGlobalUnicast() && !IsIntranetIP(ip)
}

// GetClientIP return remote IP address from request
//
func GetClientIP(req *http.Request) (remoteIP string) {
	for _, h := range []string{"X-CLIENT-IP", "X-REAL-IP", "X-FORWARDED-FOR"} {
		for _, ipToCheck := range strings.Split(req.Header.Get(h), ",") {
			// header can contain spaces too, strip those out.
			ipToCheck = strings.TrimSpace(ipToCheck)
			ip := net.ParseIP(ipToCheck)
			if ip == nil || !IsPublicIP(ip) {
				continue
			}

			return ipToCheck
		}
	}
	if strings.ContainsRune(req.RemoteAddr, ':') {
		remoteIP, _, _ = net.SplitHostPort(req.RemoteAddr)
	} else {
		remoteIP = req.RemoteAddr
	}

	return
}
