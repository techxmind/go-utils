package object

import (
	"strconv"
	"strings"
)

// GetValue get value of object by key path.
//
// obj's real type can be either map[string]interface{} or []interface{},
// otherwise return not exists.
// Usually, obj is from json universal unmarshal.
// example:
//   var obj interface{}
//   json.Unmarshal(jsonStr, &obj)
//   GetValue(obj, "path.to.node")
//
func GetValue(obj interface{}, keyPath string) (val interface{}, exists bool) {
	for _, key := range strings.Split(keyPath, ".") {
		if obj == nil {
			return nil, false
		}
		switch v := obj.(type) {
		case map[string]interface{}:
			if obj, exists = v[key]; !exists {
				return nil, false
			}
		case []interface{}:
			if i, err := strconv.Atoi(key); err != nil {
				return nil, false
			} else if i < 0 || i >= len(v) {
				return nil, false
			} else {
				obj = v[i]
			}
		default:
			return nil, false
		}
	}

	return obj, true
}
