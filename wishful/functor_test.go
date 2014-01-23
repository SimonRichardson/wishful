package wishful

import (
	"testing"
	"testing/quick"
)

// Identity

func Test_MapWithIdentity(t *testing.T) {
	f := func(v int) Id {
		return Id{v + 1}
	}
	g := func(v int) Id {
		res := Id{v}.Map(func(x AnyVal) AnyVal {
			return x.(int) + 1
		})
		return res.(Id)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_MapWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return Some{v + 1}
	}
	g := func(v int) Option {
		res := Some{v}.Map(func(x AnyVal) AnyVal {
			return x.(int) + 1
		})
		return res.(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MapWithOptionNone(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		res := None{}.Map(func(x AnyVal) AnyVal {
			return x.(int) + 1
		})
		return res.(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
