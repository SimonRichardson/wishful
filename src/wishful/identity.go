package wishful

type Identity struct {
    x AnyVal
}

func (o Identity) of(x AnyVal) Identity {
    return Identity{x}
}

// Methods
func (o Identity) chain(f func(x AnyVal) Identity) Identity {
    return f(o.x)
}

// Derived
func (o Identity) ap(x Identity) Identity {
    return o.chain(func(f AnyVal) Identity {
        return x.fmap(f.(func(f AnyVal) AnyVal))
    })
}
func (o Identity) fmap(f func(x AnyVal) AnyVal) Identity {
    return o.chain(func(x AnyVal) Identity {
        return Identity{f(x)}
    })
}