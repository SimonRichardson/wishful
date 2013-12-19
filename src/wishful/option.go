package wishful

type Value interface {}

type Option interface {
    // Methods
    Chain(func(v Value) Option) Option
    GetOrElse(x Value) Value
    OrElse(x Option) Option

    // Derived
    Map(func(v Value) Value) Option
}

type Some struct {
    x Value
}

type None struct {}

// Methods
func (o Some) Chain(f func(x Value) Option) Option {
    return f(o.x)
}
func (o None) Chain(f func(x Value) Option) Option {
    return o
}

func (o Some) OrElse(x Option) Option {
    return o
}
func (o None) OrElse(x Option) Option {
    return x
}

func (o Some) GetOrElse(x Value) Value {
    return o.x
}
func (o None) GetOrElse(x Value) Value {
    return x
}

// Derived
func (o Some) Map(f func(x Value) Value) Option {
    return o.Chain(func(x Value) Option {
        return Some{f(x)}
    })
}
func (o None) Map(f func(x Value) Value) Option {
    return o
}
