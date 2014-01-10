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
        return apply(identity)(v).(int)
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
        return compose(identity)(identity)(v).(int)
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
        return constant(v)(v + 1).(int)
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
        return identity(v).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}