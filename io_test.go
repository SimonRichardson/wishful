package wishful

import (
    "testing"
    "testing/quick"
)

func TestIOOf(t *testing.T) {
    f := func(v int) int {
        return v
    }
    g := func(v int) int {
        return IO{}.of(v).unsafePerform().(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// chain
func TestChainWithIO(t *testing.T) {
    f := func(v int) int {
        return v * 2
    }
    g := func(v int) int {
        return IO{}.of(v).chain(func (x AnyVal) IO {
            return IO{}.of(x.(int) * 2)
        }).unsafePerform().(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// ap
func TestApWithIO(t *testing.T) {
    f := func(v int) int {
        return v
    }
    g := func(v int) int {
        return IO{}.of(identity).ap(IO{}.of(v)).unsafePerform().(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// fmap
func TestFmapWithIO(t *testing.T) {
    f := func(v int) int {
        return v + 1
    }
    g := func(v int) int {
        return IO{}.of(v).fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        }).unsafePerform().(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}