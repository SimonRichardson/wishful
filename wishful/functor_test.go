package wishful

import (
	"testing"
	"testing/quick"
)

// Identity

func Test_MapWithIdentity(t *testing.T) {
	f := func(v int) Id {
		return NewId(v + 1)
	}
	g := func(v int) Id {
		res := NewId(v).Map(Inc)
		return res.(Id)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_MapWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return NewSome(v + 1)
	}
	g := func(v int) Option {
		res := NewSome(v).Map(Inc)
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
		res := None{}.Map(Inc)
		return res.(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Promise

func Test_MapWithPromise(t *testing.T) {
	f := func(v int) int {
		return v + 1
	}
	g := func(v int) int {
		pro := Promise{}.Of(v).(Promise)
		fun := pro.Map(Inc)
		p := fun.(Promise)
		return p.Fork(func(x AnyVal) AnyVal {
			return x
		}).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
