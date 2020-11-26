package compare

import (
	"strings"

	"github.com/techxmind/go-utils/itype"
)

const EPSILON float64 = 1e-9

// FloatEquals check the difference of a and b is less than epsilon
//
func FloatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

// Number compare number values
//
func Number(a, b interface{}) int {
	fa := itype.Float(a)
	fb := itype.Float(b)

	if FloatEquals(fa, fb) {
		return 0
	}

	if fa-fb > 0 {
		return 1
	} else {
		return -1
	}
}

// Object compare scalar values.
// Scalar value is number, boolean, string.
// The not-scalar value is treated as zero value.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
func Object(a, b interface{}) int {
	if a == nil && b == nil {
		return 0
	}

	ta := itype.GetType(a)
	tb := itype.GetType(b)

	if ta == itype.NUMBER || ta == itype.BOOL || tb == itype.NUMBER || tb == itype.BOOL {
		return Number(a, b)
	}

	if ta == itype.STRING || tb == itype.STRING {
		sa := ""
		sb := ""
		if ta == itype.STRING {
			sa = a.(string)
		}
		if tb == itype.STRING {
			sb = b.(string)
		}
		return strings.Compare(sa, sb)
	}

	return 0
}
