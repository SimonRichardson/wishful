package wishful

type State struct {
    run func(a AnyVal) (AnyVal, AnyVal)
}

func (o State) of(x AnyVal) State {
    return State{func(y AnyVal) (AnyVal, AnyVal) {
        return x, y
    }}
}

// Methods
func (o State) chain(f func(x AnyVal) State) State {
    return State{func(s AnyVal) (AnyVal, AnyVal) {
        a, b := o.run(s)
        return f(a).run(b)
    }}
}
func (o State) evalState(x AnyVal) AnyVal {
    a, _ := o.run(x)
    return a
}
func (o State) execState(x AnyVal) AnyVal {
    _, b := o.run(x)
    return b
}

// Derived
func (o State) ap(x State) State {
    return o.chain(func(f AnyVal) State {
        return x.fmap(f.(func(f AnyVal) AnyVal))
    })
}
func (o State) fmap(f func(x AnyVal) AnyVal) State {
    return o.chain(func(x AnyVal) State {
        return State{}.of(f(x))
    })
}