package useful

import (
	"testing"
	"testing/quick"
)

func guard(a func(x []int) int) func(x []int) int {
	return func(x []int) int {
		if len(x) < 1 {
			return 0
		}
		return a(x)
	}
}

// Manual tests
func Test_Lens(t *testing.T) {
	f := guard(func(x []int) int {
		return x[0]
	})
	g := guard(func(x []int) int {
		return Lens{}.SliceLens(0).Run(x).Get().(int)
	})
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
