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
	return GetObject(obj, keyPath, false)
}

// GetObject get child object from object.
//
// obj's real type can be either map[string]interface{} or []interface{},
// otherwise return not exists.
// If createIfNotExists is true and keyPath don't contains array access,
// will create not exists nodes in specified keyPath with type map[string]interface{}, and return exists false.
//
// Notice: keyPath = "" means obj it self
//         but keyPath = ".xxx" contains an empty path, it means obj[""]["xxx"]
//         anyway, avoid using empty path.
//
// Usually, obj is from json universal unmarshal.
// example:
//   var obj interface{}
//   json.Unmarshal(jsonStr, &obj)
//   GetObject(obj, "path.to.node", false)
//
func GetObject(obj interface{}, keyPath string, createIfNotExists bool) (val interface{}, exists bool) {
	if keyPath == "" {
		return obj, true
	}
	exists = true
	for _, key := range strings.Split(keyPath, ".") {
		switch v := obj.(type) {
		case map[string]interface{}:
			obj, exists = v[key]
			if !exists {
				if !createIfNotExists {
					return nil, false
				}
				obj = make(map[string]interface{})
				v[key] = obj
			}
		case []interface{}:
			if i, err := strconv.Atoi(key); err != nil {
				return nil, false
			} else {
				if i < 0 || i >= len(v) {
					return nil, false
				} else {
					obj = v[i]
				}
			}
		default:
			return nil, false
		}
	}

	return obj, exists
}
