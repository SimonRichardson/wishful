package useful

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_Store_Map(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		store := NewStore(Identity, func() Any {
			return x
		})
		fun := store.Map(Identity)
		return fun.(Store).Extract().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Store_Extract(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		store := NewStore(Identity, func() Any {
			return x
		})
		return store.Extract().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Store_Extend(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		store := NewStore(Identity, func() Any {
			return x
		})
		ext := store.Extend(func(x Store) Any {
			return x.Extract()
		})
		return ext.Extract().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
