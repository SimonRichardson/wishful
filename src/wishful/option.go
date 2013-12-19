package wishful

type Option interface {
    // Methods
    chain(func(v AnyVal) Option) Option
    concat(x Option) Option
    getOrElse(x AnyVal) AnyVal
    orElse(x Option) Option

    // Derived
    ap(x Option) Option
    fmap(func(v AnyVal) AnyVal) Option
}

type Some struct {
    x AnyVal
}

type None struct {}

func of(x AnyVal) Option {
    return Some{x}
}
func empty() Option {
    return None{}
}

// Methods
func (o Some) chain(f func(x AnyVal) Option) Option {
    return f(o.x)
}
func (o None) chain(f func(x AnyVal) Option) Option {
    return o
}

func (o Some) concat(x Option) Option {
    return o.chain(func(a AnyVal) Option {
        return x.fmap(func(b AnyVal) AnyVal {
            s0 := a.(Semigroup)
            s1 := b.(Semigroup)
            return s0.concat(s1)
        })
    })
}
func (o None) concat(x Option) Option {
    return o
}

func (o Some) orElse(x Option) Option {
    return o
}
func (o None) orElse(x Option) Option {
    return x
}

func (o Some) getOrElse(x AnyVal) AnyVal {
    return o.x
}
func (o None) getOrElse(x AnyVal) AnyVal {
    return x
}

// Derived
func (o Some) ap(x Option) Option {
    return o.chain(func(f AnyVal) Option {
        return x.fmap(f.(func(f AnyVal) AnyVal))
    })
}
func (o None) ap(x Option) Option {
    return o
}
func (o Some) fmap(f func(x AnyVal) AnyVal) Option {
    return o.chain(func(x AnyVal) Option {
        return Some{f(x)}
    })
}
func (o None) fmap(f func(x AnyVal) AnyVal) Option {
    return o
}
