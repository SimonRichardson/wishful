package wishful

import (
	"testing"
	"testing/quick"
)

// Identity

func Test_ChainWithIdentity(t *testing.T) {
	f := func(v int) Id {
		return Id{v + 1}
	}
	g := func(v int) Id {
		a := Id{v}.Chain(func(x AnyVal) Monad {
			return Id{x.(int) + 1}
		})
		return a.(Id)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_ChainWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return Some{v + 1}
	}
	g := func(v int) Option {
		return Some{v}.Chain(func(x AnyVal) Monad {
			return Some{x.(int) + 1}
		})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ChainWithOptionNone(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Chain(func(x AnyVal) Monad {
			return Some{x.(int) + 1}
		})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
