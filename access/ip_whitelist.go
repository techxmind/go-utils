package access

import (
	"strings"
)

type IPWhitelist struct {
	list [][]string
}

// Value格式: 111.10.11.3, 111.10.11.% 111.10.11.*
//
//   NewIPWhitelist("192.168.1.1", "192.168.1.2,192.168.1.3", []string{"10.0.0.1"}, []interface{}{"127.0.0.1"})
//
func NewIPWhitelist(whitelist ...interface{}) *IPWhitelist {
	list := make([][]string, 0, len(whitelist))
	coll := func(str string) {
		for _, subItem := range strings.Split(strings.TrimSpace(str), ",") {
			if subItem == "" {
				continue
			}
			list = append(list, strings.Split(strings.TrimSpace(subItem), "."))
		}
	}
	for _, item := range whitelist {
		switch v := item.(type) {
		case string:
			coll(v)
		case []interface{}:
			for _, elem := range v {
				if str, ok := elem.(string); ok {
					coll(str)
				}
			}
		case []string:
			for _, str := range v {
				coll(str)
			}
		}
	}
	return &IPWhitelist{
		list: list,
	}
}

func (l *IPWhitelist) Contains(ip string) bool {
	segments := strings.Split(ip, ".")

Outer:
	for _, item := range l.list {
		if len(item) > len(segments) {
			continue
		}
		for i, segment := range item {
			if segment == "*" || segment == "%" || segment == segments[i] {
				continue
			}
			continue Outer
		}
		return true
	}

	return false
}
