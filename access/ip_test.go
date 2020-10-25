package access

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIntranetIPAddress(t *testing.T) {
	ast := assert.New(t)

	ast.True(IsIntranetIPAddress("192.168.1.1"))
	ast.True(IsIntranetIPAddress("10.0.0.1"))
	ast.True(IsIntranetIPAddress("172.31.255.255"))
	ast.False(IsIntranetIPAddress("17.8.8.8"))
	ast.True(IsIntranetIPAddress("127.0.0.1"))
	ast.False(IsIntranetIPAddress("114.114.114.114"))
	ast.False(IsIntranetIPAddress(""))
	ast.False(IsIntranetIPAddress("invalid ip"))
}

func TestIsPublicIPAddress(t *testing.T) {
	ast := assert.New(t)

	ast.True(IsPublicIPAddress("114.114.114.114"))
	ast.False(IsPublicIPAddress("192.168.0.101"))
	ast.False(IsPublicIPAddress("114.114.114"))
}

func TestGetClientIP(t *testing.T) {
	ast := assert.New(t)

	req := httptest.NewRequest("GET", "http://techxmind.com", nil)

	req.RemoteAddr = "192.168.0.10:8080"
	ast.Equal("192.168.0.10", GetClientIP(req))

	req.Header.Add("X-Forwarded-For", "192.168.0.106 , 116.116.116.116")
	ast.Equal("116.116.116.116", GetClientIP(req))

	req.Header.Add("X-Real-Ip", "192.168.0.105 , 115.115.115.115")
	ast.Equal("115.115.115.115", GetClientIP(req))

	req.Header.Add("X-Client-Ip", "192.168.0.104 , 114.114.114.114")
	ast.Equal("114.114.114.114", GetClientIP(req))
}
