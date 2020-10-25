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
