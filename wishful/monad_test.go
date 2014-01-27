package wishful

import (
	"testing"
	"testing/quick"
)

// Identity

func Test_ChainWithIdentity(t *testing.T) {
	f := func(v int) Id {
		return NewId(v + 1)
	}
	g := func(v int) Id {
		a := NewId(v).Chain(func(x AnyVal) Monad {
			return NewId(Inc(x))
		})
		return a.(Id)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// IdentityT

func Test_ChainWithIdentityT(t *testing.T) {
	f := func(v int) Id {
		return NewId(Inc(v))
	}
	g := func(v int) Id {
		M := NewIdT(Id{})

		program := M.Of(v)
		mon := program.(Monad).Chain(func(x AnyVal) Monad {
			app := Id{}.Of(Inc(x))
			return app.(Monad)
		})
		return mon.(IdT).Run.(Id)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// IO

func Test_ChainWithIO(t *testing.T) {
	f := func(v int) int {
		return v + 1
	}
	g := func(v int) int {
		app := IO{}.Of(v)
		mon := app.(Monad).Chain(func(x AnyVal) Monad {
			return IO{}.Of(Inc(x)).(Monad)
		})
		io := mon.(IO)
		return io.UnsafePerform().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_ChainWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return NewSome(v + 1)
	}
	g := func(v int) Option {
		return NewSome(v).Chain(func(x AnyVal) Monad {
			return NewSome(Inc(x))
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
			return NewSome(Inc(x))
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
