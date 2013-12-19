package wishful

import (
    "testing"
    "testing/quick"
)

func TestIdentityOf(t *testing.T) {
    f := func(v int) Identity {
        return Id{v}
    }
    g := func(v int) Identity {
        return Id{1}.of(v)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// chain
func TestChainWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Id{v * 2}
    }
    g := func(v int) Identity {
        return Id{v}.chain(func (x AnyVal) Identity {
            return Id{x.(int) * 2}
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// ap
func TestApWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Id{v}
    }
    g := func(v int) Identity {
        return Id{func(x AnyVal) AnyVal {
            return x
        }}.ap(Id{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// fmap
func TestFmapWithIdentity(t *testing.T) {
    f := func(v int) Identity {
        return Id{v + 1}
    }
    g := func(v int) Identity {
        return Id{v}.fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}