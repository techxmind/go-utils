package access

import (
	"strings"
)

//  NewIPWhitelist("192.168.1.1", "192.168.1.0/24,192.168.1.3", []string{"10.0.0.1"}, []interface{}{"127.0.0.1"})
//
func NewIPWhitelist(vals ...interface{}) *IPList {

	list := make([]string, 0, len(vals)+2)

	// add localhost
	list = append(list, "127.0.0.1/8", "::1/128")

	coll := func(str string) {
		for _, subItem := range strings.Split(strings.TrimSpace(str), ",") {
			subItem = strings.TrimSpace(subItem)
			if subItem == "" {
				continue
			}
			list = append(list, subItem)
		}
	}

	for _, item := range vals {
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

	return NewIPList(list...)
}
