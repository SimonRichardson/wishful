package wishful

import (
    "testing"
    "testing/quick"
)

func TestIdentityOf(t *testing.T) {
    f := func(v int) Identity {
        return Identity{v}
    }
    g := func(v int) Identity {
        return Identity{}.of(v)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// chain
func TestChainWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Identity{v * 2}
    }
    g := func(v int) Identity {
        return Identity{v}.chain(func (x AnyVal) Identity {
            return Identity{x.(int) * 2}
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// ap
func TestApWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Identity{v}
    }
    g := func(v int) Identity {
        return Identity{identity}.ap(Identity{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// fmap
func TestFmapWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Identity{v + 1}
    }
    g := func(v int) Identity {
        return Identity{v}.fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}