package wishful

type Identity interface {
    // Methods
    chain(func(v AnyVal) Identity) Identity

    // Derived
    ap(x Identity) Identity
    fmap(func(v AnyVal) AnyVal) Identity
}

type Id struct {
    x AnyVal
}

func (o Id) of(x AnyVal) Identity {
    return Id{x}
}

// Methods
func (o Id) chain(f func(x AnyVal) Identity) Identity {
    return f(o.x)
}

// Derived
func (o Id) ap(x Identity) Identity {
    return o.chain(func(f AnyVal) Identity {
        return x.fmap(f.(func(f AnyVal) AnyVal))
    })
}
func (o Id) fmap(f func(x AnyVal) AnyVal) Identity {
    return o.chain(func(x AnyVal) Identity {
        return Id{f(x)}
    })
}