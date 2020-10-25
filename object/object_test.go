package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	ast := assert.New(t)

	obj := map[string]interface{}{
		"k1": map[string]interface{}{
			"k2": []interface{}{1, "2", 3},
		},
	}

	v, ok := GetValue(obj, "k1.k2.0")
	ast.Equal(1, v)
	ast.True(ok)

	v, ok = GetValue(obj, "k1.k2.1")
	ast.Equal("2", v)
	ast.True(ok)

	v, ok = GetValue(obj, "k1.k2.3")
	ast.Nil(v)
	ast.False(ok)
}
