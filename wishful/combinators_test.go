package wishful

import (
	"testing"
	"testing/quick"
)

// A combinator
func TestApply(t *testing.T) {
	f := func(v int) int {
		return v
	}
	g := func(v int) int {
		return Apply(Identity)(v).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// B combinator
func TestCompose(t *testing.T) {
	f := func(v int) int {
		return v
	}
	g := func(v int) int {
		return Compose(Identity)(Identity)(v).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// K combinator
func TestConstant(t *testing.T) {
	f := func(v int) int {
		return v
	}
	g := func(v int) int {
		return Constant(v)(v + 1).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// I combinator
func TestIdentity(t *testing.T) {
	f := func(v int) int {
		return v
	}
	g := func(v int) int {
		return Identity(v).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// T combinator
func TestThrush(t *testing.T) {
	f := func(v int) int {
		return v
	}
	g := func(v int) int {
		return Thrush(v)(Identity).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
