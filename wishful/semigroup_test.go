package wishful

import (
	"testing"
	"testing/quick"
)

// Create a semi-group for testing with
type IntSemigroup struct {
	x AnyVal
}

func (o IntSemigroup) Concat(x Semigroup) Semigroup {
	a := x.(IntSemigroup)
	return IntSemigroup{o.x.(int) + a.x.(int)}
}

// Identity

func Test_ConcatWithId(t *testing.T) {
	f := func(v int) Option {
		return Id{IntSemigroup{v + v}}
	}
	g := func(v int) Option {
		return Id{IntSemigroup{v}}.Concat(Id{IntSemigroup{v}})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_ConcatWithOptionSomeAndSome(t *testing.T) {
	f := func(v int) Option {
		return Some{IntSemigroup{v + v}}
	}
	g := func(v int) Option {
		return Some{IntSemigroup{v}}.Concat(Some{IntSemigroup{v}})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ConcatWithOptionSomeAndNone(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return Some{IntSemigroup{v}}.Concat(None{})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ConcatWithOptionNoneAndSome(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Concat(Some{IntSemigroup{v}})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ConcatWithOptionNoneAndNone(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Concat(None{})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
