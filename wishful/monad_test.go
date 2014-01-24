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
			return Id{Inc(x)}
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
			return Some{Inc(x)}
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
			return Some{Inc(x)}
		})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Promise

func Test_ChainWithPromise(t *testing.T) {
	f := func(v int) int {
		return v + 1
	}
	g := func(v int) int {
		pro := Promise{}.Of(v).(Promise)
		fun := pro.Chain(func(x AnyVal) Monad {
			return Promise{}.Of(Inc(x)).(Monad)
		})
		p := fun.(Promise)
		return p.Fork(func(x AnyVal) AnyVal {
			return x
		}).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
