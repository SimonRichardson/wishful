package useful

import (
	"testing"
	"testing/quick"
)

func getGuard(a func(x []int) int) func(x []int) int {
	return func(x []int) int {
		if len(x) < 1 {
			return 0
		}
		return a(x)
	}
}

func setGuard(a func(x []int, y int) []int) func(x []int, y int) []int {
	return func(x []int, y int) []int {
		if len(x) < 2 {
			return x
		}
		return a(x, y)
	}
}

// Manual tests

func Test_Lens_SliceLensGet_ReturnsCorrectValue(t *testing.T) {
	f := getGuard(func(x []int) int {
		return x[0]
	})
	g := getGuard(func(x []int) int {
		a := SliceIndex{0}
		return Lens{}.SliceLens(a).Run(x).Get().(int)
	})
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_SliceLensSet_ReturnsCorrectValue(t *testing.T) {
	f := setGuard(func(x []int, y int) []int {
		num := len(x)
		val := make([]int, num, num)
		copy(val, x)
		val[1] = y
		return val
	})
	g := setGuard(func(x []int, y int) []int {
		a := SliceIndex{1}
		return Lens{}.SliceLens(a).Run(x).Set(y).([]int)
	})
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
