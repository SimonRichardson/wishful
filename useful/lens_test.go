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

func Test_Lens_IdGet_ReturnsCorrectValue(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		return Lens{}.Id().Run(x).Get().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_IdSet_ReturnsCorrectValue(t *testing.T) {
	f := func(x int, y int) int {
		return y
	}
	g := func(x int, y int) int {
		return Lens{}.Id().Run(x).Set(y).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_SliceGet_ReturnsCorrectValue(t *testing.T) {
	f := getGuard(func(x []int) int {
		return x[0]
	})
	g := getGuard(func(x []int) int {
		a := SliceIndex{0}
		return Lens{}.Slice(a).Run(x).Get().(int)
	})
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_SliceSet_ReturnsCorrectValue(t *testing.T) {
	f := setGuard(func(x []int, y int) []int {
		num := len(x)
		val := make([]int, num, num)
		copy(val, x)
		val[1] = y
		return val
	})
	g := setGuard(func(x []int, y int) []int {
		a := SliceIndex{1}
		return Lens{}.Slice(a).Run(x).Set(y).([]int)
	})
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_IdCompose_ReturnsCorrectValue(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		return Lens{}.Id().Compose(Lens{}.Id()).Run(x).Get().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Lens_IdAndThen_ReturnsCorrectValue(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		return Lens{}.Id().AndThen(Lens{}.Id()).Run(x).Get().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
