package wishful

import (
    "testing"
    "testing/quick"
)

// Create a semi-group for testing with
type IntSemigroup struct {
    x AnyVal
}
func (o IntSemigroup) concat(x Semigroup) Semigroup {
    a := x.(IntSemigroup)
    return IntSemigroup{o.x.(int) + a.x.(int)}
}

func TestOf(t *testing.T) {
    f := func(v int) Option {
        return Some{1}
    }
    g := func(v int) Option {
        return of(1)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestEmpty(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return empty()
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// chain
func TestChainWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v * 2}
    }
    g := func(v int) Option {
        return Some{v}.chain(func (x AnyVal) Option {
            return Some{x.(int) * 2}
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestChainWithOptionNone(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.chain(func (x AnyVal) Option {
            return Some{x.(int) * 2}
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// concat
func TestConcatWithOptionSomeAndSome(t *testing.T) {
    f := func(v int) Option {
        return Some{IntSemigroup{v + v}}
    }
    g := func(v int) Option {
        return Some{IntSemigroup{v}}.concat(Some{IntSemigroup{v}})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestConcatWithOptionSomeAndNone(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return Some{IntSemigroup{v}}.concat(None{})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestConcatWithOptionNoneAndSome(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.concat(Some{IntSemigroup{v}})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestConcatWithOptionNoneAndNone(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.concat(None{})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// getOrElse
func TestGetOrElseWithOptionSome(t *testing.T) {
    f := func(v int) int {
        return v
    }
    g := func(v int) int {
        return Some{v}.getOrElse(v + 1).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestGetOrElseWithOptionNone(t *testing.T) {
    f := func(v int) int {
        return v
    }
    g := func(v int) int {
        return None{}.getOrElse(v).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// orElse
func TestOrElseWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v}
    }
    g := func(v int) Option {
        return Some{v}.orElse(Some{v + 1})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestOrElseWithOptionNone(t *testing.T) {
    f := func(v int) Option {
        return Some{v}
    }
    g := func(v int) Option {
        return None{}.orElse(Some{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// ap
func TestApWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v}
    }
    g := func(v int) Option {
        return Some{func(x AnyVal) AnyVal {
            return x
        }}.ap(Some{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestApWithOptionNoneForApMethod(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return Some{func(x AnyVal) AnyVal {
            return x
        }}.ap(None{})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestApWithOptionNoneForApConstructor(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.ap(None{})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestApWithOptionNoneForApConstructorWithSome(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.ap(Some{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// fmap
func TestFmapWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v + 1}
    }
    g := func(v int) Option {
        return Some{v}.fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestFmapWithOptionNone(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
