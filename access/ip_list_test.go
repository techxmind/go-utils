package access

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPList(t *testing.T) {
	l := NewIPList(
		"192.168.0.1",
		"8.8.8.0/24",
	)

	assert.True(t, l.Contains("192.168.0.1"))
	assert.False(t, l.Contains("192.168.0.2"))
	assert.True(t, l.Contains("8.8.8.8"))
	assert.True(t, l.Contains("8.8.8.9"))
}
