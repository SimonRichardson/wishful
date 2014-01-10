package wishful

import (
    "testing"
    "testing/quick"
)

func TestStateOf(t *testing.T) {
    f := func(v int) (int, int) {
        return v, v
    }
    g := func(v int) (int, int) {
        a, b := State{}.of(v).run(v)
        return a.(int), b.(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestStateEvalState(t *testing.T) {
    f := func(v int, w int) int {
        return v
    }
    g := func(v int, w int) int {
        return State{}.of(v).evalState(w).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}
func TestStateExecState(t *testing.T) {
    f := func(v int, w int) int {
        return w
    }
    g := func(v int, w int) int {
        return State{}.of(v).execState(w).(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// chain
func TestChainWithState(t *testing.T) {
    f := func(v int, w int) (int, int) {
        return v * 2, w
    }
    g := func(v int, w int) (int, int) {
        a, b := State{}.of(v).chain(func (x AnyVal) State {
            return State{}.of(x.(int) * 2)
        }).run(w)
        return a.(int), b.(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// ap
func TestApWithState(t *testing.T) {
    f := func(v int, w int) (int, int) {
        return v, w
    }
    g := func(v int, w int) (int, int) {
        a, b := State{}.of(identity).ap(State{}.of(v)).run(w)
        return a.(int), b.(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}

// fmap
func TestFmapWithState(t *testing.T) {
    f := func(v int, w int) (int, int) {
        return v + 1, w
    }
    g := func(v int, w int) (int, int) {
        a, b := State{}.of(v).fmap(func (x AnyVal) AnyVal {
            return x.(int) + 1
        }).run(w)
        return a.(int), b.(int)
    }
    if err := quick.CheckEqual(f, g, nil); err != nil {
        t.Error(err)
    }
}