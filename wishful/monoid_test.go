package wishful

import (
	"testing"
	"testing/quick"
)

// Option

func Test_EmptyWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return Some{v}.Empty()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EmptyWithOptionNone(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Empty()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
