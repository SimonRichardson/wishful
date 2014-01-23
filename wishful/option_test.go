package wishful

import (
	"testing"
	"testing/quick"
)

func TestOptionOf(t *testing.T) {
	f := func(v int) Option {
		return Some{v}
	}
	g := func(v int) Option {
		return Some{}.Of(v).(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
func TestOptionEmpty(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Empty().(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
