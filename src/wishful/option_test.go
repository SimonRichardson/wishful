package wishful

import (
    "testing"
    "testing/quick"
    "math/rand"
    "reflect"
)

func (o Some) Generate(rand *rand.Rand, size int) reflect.Value {
    return reflect.ValueOf(Some{rand.Intn(size) - size / 2})
}

func (o None) Generate(rand *rand.Rand, size int) reflect.Value {
    return reflect.ValueOf(None{})
}

func TestChainWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v * 2}
    }
    g := func(v int) Option {
        return Some{v}.Chain(func (x Value) Option {
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
        return None{}.Chain(func (x Value) Option {
            return Some{x.(int) * 2}
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

func TestMapWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v + 1}
    }
    g := func(v int) Option {
        return Some{v}.Map(func (x Value) Value {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

func TestMapWithOptionNone(t *testing.T) {
    f := func(v int) Option {
        return None{}
    }
    g := func(v int) Option {
        return None{}.Map(func (x Value) Value {
            return x.(int) + 1
        })
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

func TestGetOrElseWithOptionSome(t *testing.T) {
    f := func(v int) int {
        return v
    }
    g := func(v int) int {
        return Some{v}.GetOrElse(v + 1).(int)
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
        return None{}.GetOrElse(v).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

func TestOrElseWithOptionSome(t *testing.T) {
    f := func(v int) Option {
        return Some{v}
    }
    g := func(v int) Option {
        return Some{v}.OrElse(Some{v + 1})
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
        return None{}.OrElse(Some{v})
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}