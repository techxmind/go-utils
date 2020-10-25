package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	ast := assert.New(t)

	s1 := Rand(10)
	s2 := Rand(10)

	ast.Equal(10, len(s1))
	ast.Equal(10, len(s2))
	ast.NotEqual(s1, s2)
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand(10)
	}
}
