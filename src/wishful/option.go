package wishful

type Value interface {}

type Semigroup interface {
    concat(x Semigroup) Semigroup
}

type Option interface {
    // Methods
    chain(func(v Value) Option) Option
    concat(x Option) Option
    getOrElse(x Value) Value
    orElse(x Option) Option

    // Derived
    ap(x Option) Option
    fmap(func(v Value) Value) Option
}

type Some struct {
    x Value
}

type None struct {}

func of(x Value) Option {
    return Some{x}
}
func empty() Option {
    return None{}
}

// Methods
func (o Some) chain(f func(x Value) Option) Option {
    return f(o.x)
}
func (o None) chain(f func(x Value) Option) Option {
    return o
}

func (o Some) concat(x Option) Option {
    return o.chain(func(a Value) Option {
        return x.fmap(func(b Value) Value {
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

func (o Some) getOrElse(x Value) Value {
    return o.x
}
func (o None) getOrElse(x Value) Value {
    return x
}

// Derived
func (o Some) ap(x Option) Option {
    return o.chain(func(f Value) Option {
        return x.fmap(f.(func(f Value) Value))
    })
}
func (o None) ap(x Option) Option {
    return o
}
func (o Some) fmap(f func(x Value) Value) Option {
    return o.chain(func(x Value) Option {
        return Some{f(x)}
    })
}
func (o None) fmap(f func(x Value) Value) Option {
    return o
}
