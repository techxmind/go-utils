package compare

import (
	"testing"
)

func TestFloatEquals(t *testing.T) {
	tests := []struct {
		in  [2]float64
		out bool
	}{
		{
			in:  [2]float64{1.1 + 0.134, 1.234},
			out: true,
		},
		{
			in:  [2]float64{0.1 * 3, 0.15 * 2},
			out: true,
		},
		{
			in:  [2]float64{1.0000001 * 3, 3.00000031},
			out: false,
		},
	}

	for _, test := range tests {
		if test.out != FloatEquals(test.in[0], test.in[1]) {
			t.Errorf("(%f == %f) != %v", test.in[0], test.in[1], test.out)
		}
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		in  [2]interface{}
		out int
	}{
		{[2]interface{}{3, 3.0}, 0},
		{[2]interface{}{"3.3", 3.3}, 0},
		{[2]interface{}{true, 1}, 0},
		{[2]interface{}{2, true}, 1},
		{[2]interface{}{[]int{2}, 1}, -1},
		{[2]interface{}{"abc", "abd"}, 0},
	}

	for _, test := range tests {
		if test.out != Number(test.in[0], test.in[1]) {
			t.Errorf("(%v <=> %v) != %v", test.in[0], test.in[1], test.out)
		}
	}
}

func TestObject(t *testing.T) {
	tests := []struct {
		in  [2]interface{}
		out int
	}{
		{[2]interface{}{3, 3.0}, 0},
		{[2]interface{}{"3.3", 3.3}, 0},
		{[2]interface{}{true, 1}, 0},
		{[2]interface{}{2, true}, 1},
		{[2]interface{}{[]int{2}, 1}, -1},
		{[2]interface{}{"abc", "abd"}, -1},
	}

	for _, test := range tests {
		if test.out != Object(test.in[0], test.in[1]) {
			t.Errorf("(%v <=> %v) != %v", test.in[0], test.in[1], test.out)
		}
	}
}
