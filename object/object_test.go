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

func TestGetObject(t *testing.T) {
	ast := assert.New(t)
	obj := map[string]interface{}{
		"a": []interface{}{
			map[string]interface{}{
				"b": 0,
			},
		},
	}
	type checkItem struct {
		keyPath string
		create  bool
		val     interface{}
		exists  bool
	}
	type testCase struct {
		obj       interface{}
		checkList []checkItem
	}
	tests := []testCase{
		{
			obj,
			[]checkItem{
				{"a", false, obj["a"], true},
				{"a.0", false, obj["a"].([]interface{})[0], true},
				{"a.0.b", false, obj["a"].([]interface{})[0].(map[string]interface{})["b"], true},
				{"a.0.z", false, nil, false},
				{"a.0.c.d", true, map[string]interface{}{}, false},
				{"a.0.c.d", false, map[string]interface{}{}, true},
				{"a.10.b", false, nil, false},
				{"a.10.b", true, nil, false},
				{"", false, obj, true},
			},
		},
	}

	for i, test := range tests {
		for j, check := range test.checkList {
			val, exists := GetObject(test.obj, check.keyPath, check.create)
			ast.Equal(check.val, val, "case %d.%d: %s", i, j, check.keyPath)
			ast.Equal(check.exists, exists, "case %d.%d: %s", i, j, check.keyPath)
		}
	}
}
