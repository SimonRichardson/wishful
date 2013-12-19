package wishful

import (
    "testing"
    "testing/quick"
    "math/rand"
    "reflect"
)

func (o Some) generate(rand *rand.Rand, size int) reflect.Value {
    return reflect.ValueOf(Some{rand.Intn(size) - size / 2})
}

func (o None) generate(rand *rand.Rand, size int) reflect.Value {
    return reflect.ValueOf(None{})
}

func TestStuff(t *testing.T) {
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
