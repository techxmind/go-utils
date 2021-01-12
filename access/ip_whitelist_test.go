package access

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPWhitelist(t *testing.T) {
	ast := assert.New(t)

	whitelist := NewIPWhitelist(
		"",
		"192.168.11.1,192.168.11.2",
		[]string{"192.168.12.0/24", "10.0.0.0/16"},
		[]interface{}{"192.168.13.10,192.168.13.11"},
	)
	ast.True(whitelist.Contains("127.0.0.1"))
	ast.True(whitelist.Contains("192.168.11.1"))
	ast.True(whitelist.Contains("192.168.11.2"))
	ast.False(whitelist.Contains("192.168.11.3"))
	ast.True(whitelist.Contains("192.168.12.1"))
	ast.True(whitelist.Contains("10.0.0.1"))
	ast.True(whitelist.Contains("10.0.1.1"))
	ast.False(whitelist.Contains("10.1.1.1"))
	ast.True(whitelist.Contains("192.168.13.10"))
	ast.True(whitelist.Contains("192.168.13.11"))
	ast.False(whitelist.Contains("192.168.13"))
	ast.False(whitelist.Contains(""))
}
